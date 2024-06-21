package parser

type ConfigParser struct {
	parser Parser
}

func (cp *ConfigParser) Parse(dto interface{}) error {
	return cp.parser.Parse(dto)
}

func (cp *ConfigParser) PrintDefault(dto interface{}) error {
	return cp.parser.PrintDefault(dto)
}

func CreateConfigParser(parser Parser) *ConfigParser {
	return &ConfigParser{parser: parser}
}
