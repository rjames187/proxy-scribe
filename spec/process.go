package spec

import (
	"fmt"
	"reflect"
)

func buildPrimitive(kind string, example interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	result["type"] = kind
	if example != nil {
		result["example"] = example
	}
	return result
}

func convertBody(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range data {
		switch v := value.(type) {
		case string:
			result[key] = buildPrimitive("string", value)
		case int:
			result[key] = buildPrimitive("int", value)
		case bool:
			result[key] = buildPrimitive("bool", value)
		case []interface{}:
			array := make(map[string]interface{})
			array["type"] = "array"
			if len(v) > 0 {
				item := v[0]
				if _, ok := item.(map[string]interface{}); ok {
					array["items"] = convertBody(item.(map[string]interface{}))
				} else {
					itemType := reflect.TypeOf(item)
					array["items"] = buildPrimitive(itemType.String(), item)
				}
			}
			result[key] = array
		case map[string]interface{}:
			object := make(map[string]interface{})
			object["type"] = "object"
			object["properties"] = convertBody(v)
			result[key] = object
		default:
			fmt.Print("Unkown type detected: ")
			fmt.Print(value)
			result[key] = value
		}
	}

	return result
}