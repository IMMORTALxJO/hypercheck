package probe

import "fmt"

type Bool struct {
	value bool
}

func (p *Bool) Up(input *Input) (bool, string) {
	return p.value, fmt.Sprintf("boolean statement '%s' is %v", input.ToString(), p.value)
}

func (*Bool) GetType() string {
	return "Bool"
}

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}
