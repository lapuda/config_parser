package config_parser

import "io"

type Parser interface {
	Reader() io.Reader
	Parser() interface{}
}
