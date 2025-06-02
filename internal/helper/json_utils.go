package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
)

func GetFieldBytes(json []byte, path string) (interface{}, error) {
	result := gjson.GetBytes(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return nil, errors.New(errMsg)
	}

	return result.Value(), nil
}

func GetField(json string, path string) (interface{}, error) {
	result := gjson.Get(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return nil, errors.New(errMsg)
	}

	return result.Value(), nil
}

func GetResultBytes(json []byte, path string) (gjson.Result, error) {
	result := gjson.GetBytes(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return result, errors.New(errMsg)
	}

	return result, nil
}

func GetResult(json string, path string) (gjson.Result, error) {
	result := gjson.Get(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return result, errors.New(errMsg)
	}

	return result, nil
}

// StructToJSON converts any struct to a JSON string.
// Returns the JSON string and any error that occurred during marshaling.
func StructToJSON(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

// StructToPrettyJSON converts any struct to a formatted (pretty-printed) JSON string.
// Returns the formatted JSON string and any error that occurred during marshaling.
func StructToPrettyJSON(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to pretty JSON: %w", err)
	}
	return string(jsonBytes), nil
}

// JSONToStruct converts a JSON string to a struct.
// The 'v' parameter should be a pointer to the struct that will hold the unmarshaled data.
func JSONToStruct(jsonStr string, v interface{}) error {
	if jsonStr == "" {
		return fmt.Errorf("failed to unmarshal JSON to struct: %s", jsonStr)
	}

	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON to struct: %w", err)
	}
	return nil
}
