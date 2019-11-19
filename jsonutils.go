package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func GetJsonValue(src, key string) ([]byte, error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(src), &js)
	if err != nil {
		return nil, err
	}
	if js[key] == nil {
		return nil, errors.New(fmt.Sprintf("Failed to get key: %s", key))
	}
	res, err := json.Marshal(js[key])
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ParseMap(dat interface{}) map[string]interface{} {
	switch dat.(type) {
	case map[string]interface{}:
		return dat.(map[string]interface{})
	default:
		return nil
	}
}

func ParseFloat(dat interface{}) float64 {
	switch dat.(type) {
	case float64:
		return dat.(float64)
	default:
		return 0.0
	}
}
