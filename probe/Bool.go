package probe

import "fmt"

const BoolType = "Bool"

type Bool struct {
	value bool
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

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}
