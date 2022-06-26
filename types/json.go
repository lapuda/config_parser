package types

import (
	"encoding/json"
	"io"
)

type JsonParser struct {
	reader io.Reader
}

func (jp JsonParser) Parse(dto interface{}) error {
	decoder := json.NewDecoder(jp.reader)
	return decoder.Decode(dto)
}

func NewJsonConfigFile(reader io.Reader) *JsonParser {
	return &JsonParser{reader: reader}
}
