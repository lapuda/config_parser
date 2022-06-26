package main

import (
	"bytes"
	"fmt"
	"github.com/lapuda/config_parser/parser"
	"github.com/lapuda/config_parser/types"
	"os"
)

var testDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fileReader, error := os.Open("./examples/base_config.json")
	if error != nil {
		println(error.Error())
		return
	}
	jsonfile := types.NewJsonConfigFile(fileReader)
	parser1 := parser.CreateConfigParser(jsonfile)
	error = parser1.Parse(&testDto)
	if error != nil {
		println(error.Error())
		return
	}
	fmt.Println(testDto)

	stringReader := bytes.NewBuffer([]byte("{\"name\":\"test2\",\"age\":29}"))
	jsonfile2 := types.NewJsonConfigFile(stringReader)
	parser2 := parser.CreateConfigParser(jsonfile2)
	error = parser2.Parse(&testDto)
	if error != nil {
		println(error.Error())
		return
	}
	fmt.Println(testDto)
}
