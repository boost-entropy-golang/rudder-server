// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/services/debugger/destination/ (interfaces: DestinationDebugger)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/services/debugger/destination/eventDeliveryStatusUploader.go -package mock_destinationdebugger github.com/rudderlabs/rudder-server/services/debugger/destination/ DestinationDebugger
//

// Package mock_destinationdebugger is a generated GoMock package.
package mock_destinationdebugger

import (
	reflect "reflect"

	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	gomock "go.uber.org/mock/gomock"
)

// MockDestinationDebugger is a mock of DestinationDebugger interface.
type MockDestinationDebugger struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationDebuggerMockRecorder
	isgomock struct{}
}

// MockDestinationDebuggerMockRecorder is the mock recorder for MockDestinationDebugger.
type MockDestinationDebuggerMockRecorder struct {
	mock *MockDestinationDebugger
}

// NewMockDestinationDebugger creates a new mock instance.
func NewMockDestinationDebugger(ctrl *gomock.Controller) *MockDestinationDebugger {
	mock := &MockDestinationDebugger{ctrl: ctrl}
	mock.recorder = &MockDestinationDebuggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationDebugger) EXPECT() *MockDestinationDebuggerMockRecorder {
	return m.recorder
}

// HasUploadEnabled mocks base method.
func (m *MockDestinationDebugger) HasUploadEnabled(destID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUploadEnabled", destID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUploadEnabled indicates an expected call of HasUploadEnabled.
func (mr *MockDestinationDebuggerMockRecorder) HasUploadEnabled(destID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUploadEnabled", reflect.TypeOf((*MockDestinationDebugger)(nil).HasUploadEnabled), destID)
}

// RecordEventDeliveryStatus mocks base method.
func (m *MockDestinationDebugger) RecordEventDeliveryStatus(destinationID string, deliveryStatus *destinationdebugger.DeliveryStatusT) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordEventDeliveryStatus", destinationID, deliveryStatus)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RecordEventDeliveryStatus indicates an expected call of RecordEventDeliveryStatus.
func (mr *MockDestinationDebuggerMockRecorder) RecordEventDeliveryStatus(destinationID, deliveryStatus any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordEventDeliveryStatus", reflect.TypeOf((*MockDestinationDebugger)(nil).RecordEventDeliveryStatus), destinationID, deliveryStatus)
}

// Stop mocks base method.
func (m *MockDestinationDebugger) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockDestinationDebuggerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDestinationDebugger)(nil).Stop))
}
