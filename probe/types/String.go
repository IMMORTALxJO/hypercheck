package types

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

const StringType = "String"

type String struct {
	description string
	value       string
}

func (p *String) Up(input *Input) (bool, string) {
	left := p.GetValue()
	if input.Aggregator != "" {
		if input.Aggregator == "length" {
			probe := NewNumber("", uint64(len(left)), "int")
			return probe.Up(input)
		}
		return false, fmt.Sprintf("Unknown aggregator '%s' ( allowed 'length' )", input.Aggregator)
	}
	operator := input.Operator
	right := input.Value

	log.Debugf("string: check %s %s %s", left, operator, right)
	if operator == "==" || operator == "=" {
		return left == right, fmt.Sprintf("'%s' must be equal to '%s'", left, right)
	}
	if operator == "!=" {
		return left != right, fmt.Sprintf("'%s' must not be equal to '%s'", left, right)
	}
	if operator == "~" {
		return strings.Contains(left, right), fmt.Sprintf("must contain substring '%s'", right)
	}
	if operator == "~=" {
		var err error
		result, err := regexp.Match(right, []byte(left))
		if err != nil {
			return false, "regexp compile error"
		}
		return result, fmt.Sprintf("must match regexp '%s'", right)
	}
	if operator == "~!" {
		var err error
		result, err := regexp.Match(right, []byte(left))
		if err != nil {
			return false, "regexp compile error"
		}
		return !result, fmt.Sprintf("must not match regexp '%s'", right)
	}

	return false, fmt.Sprintf("unknown operator '%s'", operator)
}

func (*String) GetType() string {
	return StringType
}

func (p *String) GetValue() string {
	return p.value
}

func (p *String) GetDescription() string {
	return fmt.Sprintf("%s ( %s )", p.description, p.GetType())
}

func NewString(description string, value string) *String {
	return &String{
		description: description,
		value:       value,
	}
}
