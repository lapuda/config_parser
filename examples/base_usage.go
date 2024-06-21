package main

import (
	"fmt"
	"github.com/lapuda/config_parser/parser"
)

var testDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	err := parser.ParseJsonFile("./examples/base_config.json", &testDto)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(testDto)

	err = parser.ParseJsonString([]byte("{\"name\":\"test2\",\"age\":29}"), &testDto)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(testDto)
}
