package parser

import (
	"bytes"
	"github.com/lapuda/config_parser/types"
	"io"
	"os"
)

func NewJsonParser(reader io.Reader) *ConfigParser {
	parser := types.NewJsonConfigFile(reader)
	return CreateConfigParser(parser)
}

func ParseJsonFile(filename string, v interface{}) error {
	fileReader, err := os.Open(filename)
	if err != nil {
		return err
	}
	return NewJsonParser(fileReader).Parse(v)
}

func ParseJsonString(jsonBytes []byte, v interface{}) error {
	parser2 := NewJsonParser(bytes.NewBuffer(jsonBytes))
	return parser2.Parse(v)
}
