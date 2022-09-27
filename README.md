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
	fileReader, error := os.Open("./examples/base_config.json")
	if error != nil {
		println(error.Error())
		return
	}
	parser1 := parser.NewJsonParser(fileReader)
	error = parser1.Parse(&testDto)
	if error != nil {
		println(error.Error())
		return
	}
	fmt.Println(testDto)

	stringReader := bytes.NewBuffer([]byte("{\"name\":\"test2\",\"age\":29}"))
	parser2 := parser.NewJsonParser(stringReader)
	error = parser2.Parse(&testDto)
	if error != nil {
		println(error.Error())
		return
	}
	fmt.Println(testDto)
}

```
