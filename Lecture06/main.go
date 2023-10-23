package main

import (
	"encoding/json"
	"fmt"
)

func generateStruct(jsonStr string) string {
	var data map[string]string
	// json в map
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return ""
	}
	return convertToGoStruct(data, "Person")
}

func convertToGoStruct(data map[string]string, structName string) string {
	structCode := "type " + structName + " struct {\n"

	for key, fieldType := range data {
		structCode += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", key, fieldType, key)
	}

	structCode += "}\n"
	return structCode
}

func main() {
	jsonInput := `{
		"name": "string",
		"age": "int"
	}`
	fmt.Println(generateStruct(jsonInput))
}
