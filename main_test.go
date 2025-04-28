package main_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mrangel-jr/ontraport/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func returnFloat64(value int) float64 {
	return float64(value)
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected map[string]interface{}
	}{
		{
			name:  "Simple JSON",
			input: []byte(`{"one": 2}`),
			expected: map[string]interface{}{
				"one": returnFloat64(2),
			},
		},
		{
			name:  "Nested JSON",
			input: []byte(`{"one": {"two": 3}}`),
			expected: map[string]interface{}{
				"one/two": returnFloat64(3),
			},
		},
		{
			name:  "Nested JSON with array",
			input: []byte(`{"one": {"two": 3, "four": [5,6,7]}}`),
			expected: map[string]interface{}{
				"one/four/0": returnFloat64(5),
				"one/four/1": returnFloat64(6),
				"one/four/2": returnFloat64(7),
				"one/two":    returnFloat64(3),
			},
		},
		{
			name:  "Nested JSON with multiple levels",
			input: []byte(`{"one": {"two": 3, "four": [5,6,7]}, "eight": {"nine": {"ten": 11}}}`),
			expected: map[string]interface{}{
				"eight/nine/ten": returnFloat64(11),
				"one/four/0":     returnFloat64(5),
				"one/four/1":     returnFloat64(6),
				"one/four/2":     returnFloat64(7),
				"one/two":        returnFloat64(3),
			},
		},
		{
			name:     "Empty JSON",
			input:    []byte(`{}`),
			expected: map[string]interface{}{},
		},
		{
			name:     "Null JSON",
			input:    []byte(`null`),
			expected: map[string]interface{}{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := utils.UnmarshalJSON(test.input)

			require.NoError(t, err)
			for key, value := range result {
				assert.Equal(t, value, test.expected[key])
			}
			for key, value := range test.expected {
				assert.Equal(t, value, result[key])
			}
		})
	}

}

func jsonEqual(a, b []byte) (bool, error) {
	var obj1 interface{}
	var obj2 interface{}

	// Unmarshal both JSON into generic interfaces (maps, slices)
	if err := json.Unmarshal(a, &obj1); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &obj2); err != nil {
		return false, err
	}

	// Compare the unmarshaled objects
	return reflect.DeepEqual(obj1, obj2), nil
}

func TestMarshalJSON(t *testing.T) {
	// Test case for marshalling a map to JSON
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected []byte
	}{
		{
			name:     "Empty map",
			input:    map[string]interface{}{},
			expected: []byte(`{}`),
		},
		{
			name: "Simple map",
			input: map[string]interface{}{
				"one": returnFloat64(2),
			},
			expected: []byte(`{
				"one": 2
			}`),
		},
		{
			name: "Nested map",
			input: map[string]interface{}{
				"one/two": returnFloat64(3),
			},
			expected: []byte(`{
				"one": {
					"two": 3
				}
			}`),
		},
		{
			name: "Nested map with other keys",
			input: map[string]interface{}{
				"one/two":        returnFloat64(3),
				"eight/nine/ten": returnFloat64(11),
			},
			expected: []byte(`{
				"eight": {
					"nine": {
						"ten": 11
					}
				},
				"one": {
					"two": 3
				}
			}`),
		},
		{
			name: "Nested map with array",
			input: map[string]interface{}{
				"one/four/0":  returnFloat64(5),
				"one/four/1":  returnFloat64(6),
				"one/four/2":  returnFloat64(7),
				"one/two":     returnFloat64(3),
				"one/eight/0": returnFloat64(9),
				"one/eight/1": returnFloat64(10),
			},
			expected: []byte(`{
				"one": {
					"four": [5, 6, 7],
					"two": 3,
					"eight": [9, 10]
				}
			}`),
		},
		{
			name: "Complex nested map",
			input: map[string]interface{}{
				"eight/nine/ten": returnFloat64(11),
				"one/four/0":     returnFloat64(5),
				"one/four/1":     returnFloat64(6),
				"one/four/2":     returnFloat64(7),
				"one/two":        returnFloat64(3),
			},
			expected: []byte(`{
				"eight": {
					"nine": {
						"ten": 11
					}
				},
				"one": {
					"four": [5, 6, 7],
					"two": 3
				}
			}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := utils.MarshalJSON(test.input)

			require.NoError(t, err)

			equal, err := jsonEqual(result, test.expected)
			require.NoError(t, err)
			require.True(t, equal, "Expected JSON does not match the result JSON")
		})
	}
}
