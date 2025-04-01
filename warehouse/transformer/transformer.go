package transformer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/processor/types"
	"github.com/rudderlabs/rudder-server/utils/misc"

	"github.com/rudderlabs/rudder-server/utils/timeutil"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/transformer/internal/response"
	"github.com/rudderlabs/rudder-server/warehouse/transformer/internal/utils"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	violationErrors     = "violationErrors"
	redshiftStringLimit = 512
)

// Compile-time check to ensure ValidationError struct remains unchanged.
var _ = struct {
	Type     string
	Message  string
	Meta     map[string]string
	Property string
}(types.ValidationError{})

func New(conf *config.Config, logger logger.Logger, statsFactory stats.Stats) *Transformer {
	t := &Transformer{
		logger:         logger.Child("warehouse-transformer"),
		statsFactory:   statsFactory,
		now:            timeutil.Now,
		uuidGenerator:  uuid.NewString,
		loggedFileName: generateLogFileName(),
	}

	t.stats.matchedEvents = t.statsFactory.NewStat("warehouse_dest_transform_matched_events", stats.HistogramType)
	t.stats.mismatchedEvents = t.statsFactory.NewStat("warehouse_dest_transform_mismatched_events", stats.HistogramType)
	t.stats.comparisonTime = t.statsFactory.NewStat("warehouse_dest_transform_comparison_time", stats.TimerType)

	t.config.enableIDResolution = conf.GetReloadableBoolVar(false, "Warehouse.enableIDResolution")
	t.config.populateSrcDestInfoInContext = conf.GetReloadableBoolVar(true, "WH_POPULATE_SRC_DEST_INFO_IN_CONTEXT")
	t.config.maxColumnsInEvent = conf.GetReloadableIntVar(200, 1, "WH_MAX_COLUMNS_IN_EVENT")
	t.config.maxLoggedEvents = conf.GetReloadableIntVar(10000, 1, "Warehouse.maxLoggedEvents")
	t.config.concurrentTransformations = conf.GetReloadableIntVar(10, 1, "Warehouse.concurrentTransformations")
	return t
}

func (t *Transformer) Transform(_ context.Context, clientEvents []types.TransformerEvent) (res types.Response) {
	if len(clientEvents) == 0 {
		return
	}

	startTime := t.now()
	metadata := clientEvents[0].Metadata
	c := &cache{}

	defer func() {
		tags := stats.Tags{
			"workspaceId":     metadata.WorkspaceID,
			"sourceId":        metadata.SourceID,
			"sourceType":      metadata.SourceType,
			"destinationId":   metadata.DestinationID,
			"destinationType": metadata.DestinationType,
		}

		t.statsFactory.NewTaggedStat("warehouse_dest_transform_request_latency", stats.TimerType, tags).Since(startTime)
		t.statsFactory.NewTaggedStat("warehouse_dest_transform_requests", stats.CountType, tags).Increment()
		t.statsFactory.NewTaggedStat("warehouse_dest_transform_input_events", stats.HistogramType, tags).Observe(float64(len(clientEvents)))
		t.statsFactory.NewTaggedStat("warehouse_dest_transform_output_events", stats.HistogramType, tags).Observe(float64(len(res.Events)))
		t.statsFactory.NewTaggedStat("warehouse_dest_transform_output_failed_events", stats.HistogramType, tags).Observe(float64(len(res.FailedEvents)))
	}()

	results := make(chan types.TransformerResponse, len(clientEvents))
	done := make(chan struct{})

	g := errgroup.Group{}
	g.SetLimit(t.config.concurrentTransformations.Load())

	go func() {
		defer close(done)
		for r := range results {
			if r.Error != "" {
				res.FailedEvents = append(res.FailedEvents, r)
			} else {
				res.Events = append(res.Events, r)
			}
		}
	}()

	for i := range clientEvents {
		event := clientEvents[i]
		event.Metadata.SourceDefinitionType = "" // TODO: Currently, it's getting ignored during JSON marshalling Remove this once we start using it.

		g.Go(func() error {
			r, err := t.processWarehouseMessage(c, &event)
			if err != nil {
				results <- transformerResponseFromErr(&event, err)
				return nil
			}
			for _, item := range r {
				results <- types.TransformerResponse{
					Output:     item,
					Metadata:   event.Metadata,
					StatusCode: http.StatusOK,
				}
			}
			return nil
		})
	}

	_ = g.Wait()
	close(results)
	<-done
	return
}

func (t *Transformer) processWarehouseMessage(cache *cache, event *types.TransformerEvent) ([]map[string]any, error) {
	if err := t.enhanceContextWithSourceDestInfo(event); err != nil {
		return nil, fmt.Errorf("enhancing context with source and destination info: %w", err)
	}
	t.addMandatoryFields(event)
	return t.handleEvent(event, cache)
}

func (t *Transformer) enhanceContextWithSourceDestInfo(event *types.TransformerEvent) error {
	if !t.config.populateSrcDestInfoInContext.Load() {
		return nil
	}

	messageContext, ok := event.Message["context"]
	if !ok || messageContext == nil {
		messageContext = map[string]any{}
	}
	messageContextMap, ok := messageContext.(map[string]any)
	if !ok {
		return response.ErrContextNotMap
	}
	messageContextMap["sourceId"] = event.Metadata.SourceID
	messageContextMap["sourceType"] = event.Metadata.SourceType
	messageContextMap["destinationId"] = event.Metadata.DestinationID
	messageContextMap["destinationType"] = event.Metadata.DestinationType

	event.Message["context"] = messageContextMap
	return nil
}

func (t *Transformer) addMandatoryFields(event *types.TransformerEvent) {
	t.ensureMessageID(event)
	t.ensureReceivedAt(event)
}

func (t *Transformer) ensureMessageID(event *types.TransformerEvent) {
	messageID, exists := event.Message["messageId"]
	if !exists || utils.IsBlank(messageID) {
		event.Metadata.MessageID = "auto-" + t.uuidGenerator()
		event.Message["messageId"] = event.Metadata.MessageID
	}
}

func (t *Transformer) ensureReceivedAt(event *types.TransformerEvent) {
	receivedAt, exists := event.Message["receivedAt"]
	if !exists || utils.IsBlank(receivedAt) {
		t.setDefaultReceivedAt(event)
		return
	}

	strReceivedAt, isString := receivedAt.(string)
	if !isString || !utils.ValidTimestamp(strReceivedAt) {
		t.setDefaultReceivedAt(event)
	}
}

func (t *Transformer) setDefaultReceivedAt(event *types.TransformerEvent) {
	if !utils.ValidTimestamp(event.Metadata.ReceivedAt) {
		event.Metadata.ReceivedAt = t.now().Format(misc.RFC3339Milli)
	}
	event.Message["receivedAt"] = event.Metadata.ReceivedAt
}

func (t *Transformer) handleEvent(event *types.TransformerEvent, cache *cache) ([]map[string]any, error) {
	intrOpts := extractIntrOpts(event.Metadata.DestinationType, event.Message)
	destOpts := extractDestOpts(event.Metadata.DestinationType, event.Destination.Config)
	jsonPathsInfo := extractJSONPathInfo(append(intrOpts.jsonPaths, destOpts.jsonPaths...))

	eventType := strings.ToLower(event.Metadata.EventType)

	tec := &transformEventContext{
		event:         event,
		intrOpts:      &intrOpts,
		destOpts:      &destOpts,
		jsonPathsInfo: &jsonPathsInfo,
		cache:         cache,
	}

	switch eventType {
	case "track":
		return t.trackEvents(tec)
	case "identify":
		return t.identifyEvents(tec)
	case "page":
		return t.pageEvents(tec)
	case "screen":
		return t.screenEvents(tec)
	case "alias":
		return t.aliasEvents(tec)
	case "group":
		return t.groupEvents(tec)
	case "extract":
		return t.extractEvents(tec)
	case "merge":
		return t.mergeEvents(tec)
	default:
		return nil, response.NewTransformerError(fmt.Sprintf("Unknown event type: %q", eventType), http.StatusBadRequest)
	}
}

func transformerResponseFromErr(event *types.TransformerEvent, err error) types.TransformerResponse {
	var te *response.TransformerError
	if ok := errors.As(err, &te); ok {
		return types.TransformerResponse{
			Output:     nil,
			Metadata:   event.Metadata,
			Error:      te.Error(),
			StatusCode: te.StatusCode(),
		}
	}
	return types.TransformerResponse{
		Output:     nil,
		Metadata:   event.Metadata,
		Error:      response.ErrInternalServer.Error(),
		StatusCode: response.ErrInternalServer.StatusCode(),
	}
}

func (t *Transformer) getColumns(
	destType string,
	data map[string]any, metadata map[string]string,
) (map[string]any, error) {
	columns := make(map[string]any, len(data))

	// uuid_ts and loaded_at datatypes are passed from here to create appropriate columns.
	// Corresponding values are inserted when loading into the warehouse
	uuidTS := whutils.ToProviderCase(destType, "uuid_ts")
	columns[uuidTS] = model.DateTimeDataType

	if destType == whutils.BQ {
		columns["loaded_at"] = model.DateTimeDataType
	}

	for key, value := range data {
		if dataType, ok := metadata[key]; ok {
			columns[key] = dataType
		} else {
			columns[key] = dataTypeFor(destType, key, value, false)
		}
	}
	if len(columns) > t.config.maxColumnsInEvent.Load() && !utils.IsRudderSources(data) && !utils.IsDataLake(destType) {
		return nil, response.NewTransformerError(fmt.Sprintf("%s transformer: Too many columns outputted from the event", strings.ToLower(destType)), http.StatusBadRequest)
	}
	return columns, nil
}
