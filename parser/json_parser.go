package parser

import (
	"github.com/lapuda/config_parser/types"
	"io"
)

func NewJsonParser(reader io.Reader) *ConfigParser {
	parser := types.NewJsonConfigFile(reader)
	return CreateConfigParser(parser)
}
