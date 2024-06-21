package types

import (
	"encoding/json"
	"io"
)

type JsonParser struct {
	reader io.Reader
	writer io.Writer
}

func (jp JsonParser) Parse(dto interface{}) error {
	decoder := json.NewDecoder(jp.reader)
	return decoder.Decode(dto)
}

func (jp JsonParser) PrintDefault(dto interface{}) error {
	data, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = jp.writer.Write(data)
	return err
}

func NewJsonConfigFile(reader io.Reader, writer io.Writer) *JsonParser {
	return &JsonParser{reader: reader, writer: writer}
}
