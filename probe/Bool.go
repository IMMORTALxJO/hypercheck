package probe

type Bool struct {
	value bool
}

func (p *Bool) Up(input *Input) (bool, string) {
	return p.value, ""
}

func (p *Bool) GetType() string {
	return "Bool"
}

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}
