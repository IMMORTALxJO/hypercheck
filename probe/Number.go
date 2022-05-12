package probe

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Number struct {
	value  uint64
	parser Parser
}

func (p *Number) Up(input *Input) (bool, string) {
	right, err := p.parser.Parse(input.Value)
	if err != "" {
		return false, fmt.Sprintf("failed to parse value '%s'", input.ToString())
	}
	left := p.GetValue()
	operator := input.Operator
	log.Debugf("number: check %d %s %d", left, operator, right)
	if operator == "==" || operator == "=" {
		return left == right, fmt.Sprintf("%d == %d", left, right)
	}
	if operator == ">" {
		return left > right, fmt.Sprintf("%d > %d", left, right)
	}
	if operator == ">=" {
		return left >= right, fmt.Sprintf("%d >= %d", left, right)
	}
	if operator == "<" {
		return left < right, fmt.Sprintf("%d < %d", left, right)
	}
	if operator == "<=" {
		return left <= right, fmt.Sprintf("%d <= %d", left, right)
	}
	if operator == "!=" {
		return left != right, fmt.Sprintf("%d != %d", left, right)
	}
	return false, fmt.Sprintf("unknown operator '%s'", operator)
}

func (*Number) GetType() string {
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
