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

```
