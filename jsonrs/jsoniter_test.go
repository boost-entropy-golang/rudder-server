package jsonrs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsoniterJSON(t *testing.T) {
	jsonrs := &jsoniterJSON{}

	t.Run("json.RawMessage", func(t *testing.T) {
		t.Run("marshal nil", func(t *testing.T) {
			type test struct {
				A json.RawMessage `json:"a"`
			}
			data, err := jsonrs.Marshal(test{})
			require.NoError(t, err)
			require.Equal(t, `{"a":null}`, string(data), `it marshalls a nil json.RawMessage to "null"`)
		})

		t.Run("marshal invalid", func(t *testing.T) { // unique behaviour: all other implementations return an error
			type test struct {
				A json.RawMessage `json:"a"`
			}
			data, err := jsonrs.Marshal(test{A: json.RawMessage(`{someInvalid:"}`)})
			require.NoError(t, err)
			require.Equal(t, `{"a":null}`, string(data), `it marshalls an invalid json.RawMessage to "null"`)
		})

		t.Run("unmarshal null", func(t *testing.T) { // unique behaviour: all other implementations return a "null" json.RawMessage
			type test struct {
				A json.RawMessage `json:"a"`
			}
			var v test
			err := jsonrs.Unmarshal([]byte(`{"a":null}`), &v)
			require.NoError(t, err)
			require.Nil(t, v.A, `it unmarshalls a "null" value to a nil json.RawMessage`)
		})

		t.Run("unmarshal missing", func(t *testing.T) {
			type test struct {
				A json.RawMessage `json:"a"`
			}
			var v test
			err := jsonrs.Unmarshal([]byte(`{}`), &v)
			require.NoError(t, err)
			require.Nil(t, v.A, `it unmarshalls a missing value to a nil json.RawMessage`)
		})

		t.Run("unmarshal invalid", func(t *testing.T) {
			type test struct {
				A json.RawMessage `json:"a"`
			}
			var v test
			err := jsonrs.Unmarshal([]byte(`{"a":{someInvalid:"}}`), &v)
			require.Error(t, err, `it returns an error when trying to unmarshal an invalid value to json.RawMessage`)
		})
	})
}
