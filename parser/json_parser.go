package parser

import (
	"bytes"
	"github.com/lapuda/config_parser/types"
	"io"
	"os"
)

func newJsonParser(reader io.Reader) *ConfigParser {
	parser := types.NewJsonConfigFile(reader, nil)
	return CreateConfigParser(parser)
}

func newJsonWriter(writer io.Writer) *ConfigParser {
	parser := types.NewJsonConfigFile(nil, writer)
	return CreateConfigParser(parser)
}

func ParseJsonFile(filename string, v interface{}) error {
	fileReader, err := os.Open(filename)
	if err != nil {
		return err
	}
	return newJsonParser(fileReader).Parse(v)
}

func ParseJsonString(jsonBytes []byte, v interface{}) error {
	parser2 := newJsonParser(bytes.NewBuffer(jsonBytes))
	return parser2.Parse(v)
}

func PrintInterfaceDefault(v interface{}, fileName string) error {
	fileWriter, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	return newJsonWriter(fileWriter).PrintDefault(v)
}
