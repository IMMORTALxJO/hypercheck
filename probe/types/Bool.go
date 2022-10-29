package types

import "fmt"

const BoolType = "Bool"

type Bool struct {
	description string
	value       bool
}

func (p *Bool) Up(input *Input) (bool, string) {
	return p.GetValue(), fmt.Sprintf("boolean statement '%s' is %v", input.ToString(), p.value)
}

func (*Bool) GetType() string {
	return BoolType
}

func (p *Bool) GetValue() bool {
	return p.value
}

func (p *Bool) GetDescription() string {
	return fmt.Sprintf("%s ( %s )", p.description, p.GetType())
}

func NewBool(description string, value bool) *Bool {
	return &Bool{
		description: description,
		value:       value,
	}
}
