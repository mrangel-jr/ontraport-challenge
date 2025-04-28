package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

func MarshalJSON(m CustomMap) ([]byte, error) {

	if m == nil {
		return []byte("null"), nil
	}
	result := make(CustomMap)
	for key, value := range m {
		sliceKey := strings.Split(key, "/")
		result = mapJSON(result, sliceKey, value)
	}
	// Marshal the result map to JSON
	jsonData, err := json.Marshal(result)
	return jsonData, err
}

func mapJSON(result CustomMap, arrkeys []string, value interface{}) CustomMap {
	if len(arrkeys) == 0 {
		return nil
	} else if len(arrkeys) == 1 {
		result[arrkeys[0]] = value
		return result
	} else {
		key := arrkeys[0]
		nextKey := arrkeys[1]
		if _, err := strconv.Atoi(nextKey); err == nil {
			if _, ok := result[key].([]interface{}); ok {
				result[key] = append(result[key].([]interface{}), value)
			} else {
				slice := make([]interface{}, 0)
				slice = append(slice, value)
				result[key] = slice
			}
		} else {
			if _, ok := result[key]; !ok {
				result[key] = make(CustomMap)
			}
			result[key] = mapJSON(result[key].(CustomMap), arrkeys[1:], value)
		}
	}
	return result
}
