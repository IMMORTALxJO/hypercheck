package probe

import (
	"regexp"
	"strings"
)

type String struct {
	value string
}

func (p *String) Up(input *Input) (bool, string) {
	if input.Aggregator == "length" {
		probe := NewNumber(uint64(len(p.value)), "int")
		return probe.Up(input)
	}

	operator := input.Operator
	target := input.Value
	if operator == "==" || operator == "=" {
		return p.value == target, ""
	}
	if operator == "!=" {
		return p.value != target, ""
	}
	if operator == "~" {
		return strings.Contains(p.value, target), ""
	}
	if operator == "~=" {
		var err error
		result, err := regexp.Match(target, []byte(p.value))
		if err != nil {
			return false, "regexp compile error"
		}
		return result, ""
	}
	if operator == "~!" {
		var err error
		result, err := regexp.Match(target, []byte(p.value))
		if err != nil {
			return false, "regexp compile error"
		}
		return !result, ""
	}

	return false, "Unknown operator"
}

func (p *String) GetType() string {
	return "String"
}

func NewString(value string) *String {
	return &String{
		value: value,
	}
}
