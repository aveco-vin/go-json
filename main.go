// How to make use of Go JSON
// using a struct and dynamically adding fields with different data

package main

import (
	"encoding/json"
	"fmt"
)

type JSONResponse struct {
	Fields map[string]interface{}
}

func main() {
	var resp JSONResponse
	resp.Fields = make(map[string]interface{})
	resp.Fields["id"] = "test"
	resp.Fields["content"] = 123
	data1, _ := json.Marshal(resp)
	fmt.Println(json.Marshal(resp))
	fmt.Printf("%s\n", data1)

	testMap := make(map[string]interface{})
	testMap["id"] = "test"
	data2, _ := json.Marshal(testMap)
	fmt.Println(json.Marshal(testMap))
	fmt.Printf("%s", data2)
}
