package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// StructToMap parses a struct to a map
func StructToMap(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		return nil
	}

	objType := objValue.Type()
	data := make(map[string]string)
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objType.Field(i).Tag.Get("json")
		data[fieldName] = Stringify(field.Interface())
	}

	return data
}

func Stringify(value interface{}) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	// Add more cases for other types as needed
	default:
		return fmt.Sprintf("%v", v)
	}
}
