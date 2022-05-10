package probe

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type String struct {
	value string
}

func (p *String) Up(input *Input) (bool, string) {
	left := p.GetValue()
	if input.Aggregator == "length" {
		probe := NewNumber(uint64(len(left)), "int")
		return probe.Up(input)
	}
	operator := input.Operator
	right := input.Value

	log.Debugf("string: check %s %s %s", left, operator, right)
	if operator == "==" || operator == "=" {
		return left == right, fmt.Sprintf("must be equal to '%s'", right)
	}
	if operator == "!=" {
		return left != right, fmt.Sprintf("must not be equal to '%s'", right)
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

func (p *String) GetType() string {
	return "String"
}

func (p *String) GetValue() string {
	return p.value
}

func NewString(value string) *String {
	return &String{
		value: value,
	}
}
