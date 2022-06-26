package tests

import (
	"bytes"
	"errors"
	"github.com/lapuda/config_parser/parser"
	"github.com/lapuda/config_parser/types"
	"os"
	"strconv"
	"testing"
)

var dto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJSONFile(t *testing.T) {

	fileReader, error := os.Open("./base_config.json")
	if error != nil {
		println(error.Error())
		return
	}
	jsonfile := types.NewJsonConfigFile(fileReader)
	parser1 := parser.CreateConfigParser(jsonfile)
	error = parser1.Parse(&dto)
	if error != nil {
		t.Error(error)
		return
	}

	if dto.Name != "JSONFileTest" {
		t.Error(errors.New("parse json file error Name:expected:JSONFileTest，actual：" + dto.Name))
		return
	}

	if dto.Age != 28 {
		t.Error(errors.New("parse json file error, Age:expected 28，actual：" + strconv.Itoa(dto.Age)))
		return
	}
}

func TestJsonString(t *testing.T) {
	jsonfile := types.NewJsonConfigFile(bytes.NewBuffer([]byte("{\"name\":\"JsonStringTest\",\"age\":28}")))
	parser1 := parser.CreateConfigParser(jsonfile)
	error := parser1.Parse(&dto)
	if error != nil {
		t.Error(error)
		return
	}

	if dto.Name != "JsonStringTest" {
		t.Error(errors.New("parse json file error Name:expected: JsonStringTest，actual：" + dto.Name))
		return
	}

	if dto.Age != 28 {
		t.Error(errors.New("parse json file error, Age:expected 28，actual：" + strconv.Itoa(dto.Age)))
		return
	}
}
