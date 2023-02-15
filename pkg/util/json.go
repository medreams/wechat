package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// struct 转 map
func ConvertToMap(content interface{}) map[string]interface{} {
	b, err := json.Marshal(content)
	if err != nil {
		fmt.Println("ConvertToMap Marshal Error ", err)
		return nil
	}
	var result map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		fmt.Println("ConvertToMap Decode Error ", err)
		return nil
	}
	return result
}
