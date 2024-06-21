package parser

type Parser interface {
	Parse(dto interface{}) error
	PrintDefault(dto interface{}) error
}
