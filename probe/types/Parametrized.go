package types

const ParametrizedType = "Parametrized"

type Parametrized struct {
	probe Probe
	input *Input
}

func (p *Parametrized) Up(_ *Input) (bool, string) {
	if p.probe.GetType() == GeneratorType {
		return p.probe.(*Generator).Up(p.input)
	}
	if p.probe.GetType() == NumberType {
		return p.probe.(*Number).Up(p.input)
	}
	if p.probe.GetType() == ListType {
		return p.probe.(*List).Up(p.input)
	}
	if p.probe.GetType() == MapType {
		return p.probe.(*Map).Up(p.input)
	}
	if p.probe.GetType() == StringType {
		return p.probe.(*String).Up(p.input)
	}
	return p.probe.(*Bool).Up(p.input)
}

func (*Parametrized) GetType() string {
	return ParametrizedType
}

func (p *Parametrized) GetValue() Probe {
	return p.probe
}

func (p *Parametrized) GetDescription() string {
	return p.probe.GetDescription()
}

func NewParametrized(probe Probe, input *Input) *Parametrized {
	return &Parametrized{
		probe: probe,
		input: input,
	}
}
