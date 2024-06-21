# config_parser
a simple config parser!
## useage
```
  go get github.com/lapuda/config_parser
```
```
package main

import (
	"bytes"
	"fmt"
	"github.com/lapuda/config_parser/parser"
	"os"
)

var testDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// read config from json
	err := parser.ParseJsonFile("./examples/base_config.json", &testDto)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(testDto)

	// read config from json buffer
	err = parser.ParseJsonString([]byte("{\"name\":\"test2\",\"age\":292,\"sbv\":10}"), &testDto)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(testDto)

	// write config to json file
	err = parser.PrintInterfaceDefault(testDto, "config1.json")
	if err != nil {
		println(err.Error())
		return
	}
}

```
