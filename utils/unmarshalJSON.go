package utils

import (
	"encoding/json"
	"fmt"
)

type CustomMap map[string]interface{}

func UnmarshalJSON(input []byte) (CustomMap, error) {
	var data CustomMap

	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	result := make(CustomMap)
	// Example of how to access the data
	for key, value := range data {
		callMapIdentification(value, &result, key)
	}
	return result, nil
}

func callMapIdentification(value interface{}, result *CustomMap, key string) {
	if nestedCustomMap, ok := value.(map[string]interface{}); ok {
		nestedMap(nestedCustomMap, result, key)
	} else {
		setValue(key, value, result)
	}
}

func nestedMap(nestedCustomMap CustomMap, result *CustomMap, keyResult string) {
	for key, value := range nestedCustomMap {
		key = fmt.Sprintf("%s/%s", keyResult, key)
		callMapIdentification(value, result, key)
	}
}

func setValue(key string, value interface{}, result *CustomMap) {
	switch value := value.(type) {
	case string:
		(*result)[key] = value
	case float64:
		(*result)[key] = value
	case []interface{}:
		for i, v := range value {
			// Check if the value is a map
			key := fmt.Sprintf("%s/%d", key, i)
			callMapIdentification(v, result, key)
		}
	default:
		(*result)[key] = value
	}
}

func (m CustomMap) String() string {
	if m == nil {
		return "nil"
	}
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(b)
}
