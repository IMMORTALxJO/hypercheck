package probe

type Number struct {
	value  uint64
	parser Parser
}

func (p *Number) Up(input *Input) (bool, string) {
	right, err := p.parser.Parse(input.Value)
	if err != "" {
		return false, "Couldn't parse input"
	}
	left := p.value
	operator := input.Operator
	if operator == "==" || operator == "=" {
		return left == right, ""
	}
	if operator == ">" {
		return left > right, ""
	}
	if operator == ">=" {
		return left >= right, ""
	}
	if operator == "<" {
		return left < right, ""
	}
	if operator == "<=" {
		return left <= right, ""
	}
	if operator == "!=" {
		return left != right, ""
	}
	return false, "Unknown operator"
}

func (p *Number) GetType() string {
	return "Number"
}

func (p *Number) GetValue() uint64 {
	return p.value
}

func (p *Number) GetParserType() string {
	return p.parser.GetType()
}

func NewNumber(value uint64, parserName string) *Number {
	var parser Parser
	if parserName == "bytes" {
		parser = &ParserBytes{}
	} else {
		parser = &ParserInt{}
	}
	return &Number{
		value:  value,
		parser: parser,
	}
}
