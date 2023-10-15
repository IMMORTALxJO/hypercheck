package types

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const NumberType = "Number"

type Number struct {
	description string
	value       uint64
	parser      Parser
}

func (p *Number) Up(input *Input) (bool, string) {
	right, err := p.parser.Parse(input.Value)
	if err != "" {
		return false, fmt.Sprintf("failed to parse value '%s'", input.ToString())
	}
	left := p.GetValue()
	operator := input.Operator
	inputString := input.ToString()
	log.Debugf("number: check %d %s %d", left, operator, right)
	if operator == "==" || operator == "=" {
		return left == right, fmt.Sprintf("'%s' %d == %d", inputString, left, right)
	}
	if operator == ">" {
		return left > right, fmt.Sprintf("'%s' %d > %d", inputString, left, right)
	}
	if operator == ">=" {
		return left >= right, fmt.Sprintf("'%s' %d >= %d", inputString, left, right)
	}
	if operator == "<" {
		return left < right, fmt.Sprintf("'%s' %d < %d", inputString, left, right)
	}
	if operator == "<=" {
		return left <= right, fmt.Sprintf("'%s' %d <= %d", inputString, left, right)
	}
	if operator == "!=" {
		return left != right, fmt.Sprintf("'%s' %d != %d", inputString, left, right)
	}
	return false, fmt.Sprintf("unknown operator '%s'", operator)
}

func (*Number) GetType() string {
	return NumberType
}

func (p *Number) GetValue() uint64 {
	return p.value
}

func (p *Number) GetParserType() string {
	return p.parser.GetType()
}

func (p *Number) GetDescription() string {
	return fmt.Sprintf("%s ( %s )", p.description, p.GetType())
}

func NewNumber(description string, value uint64, parserName string) *Number {
	var parser Parser
	switch parserName {
	case "bytes":
		parser = &ParserBytes{}
	case "duration":
		parser = &ParserDuration{}
	default:
		parser = &ParserInt{}
	}
	return &Number{
		description: description,
		value:       value,
		parser:      parser,
	}
}
