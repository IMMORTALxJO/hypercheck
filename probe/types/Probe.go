package types

type Probe interface {
	Up(*Input) (bool, string)
	GetType() string
	GetDescription() string
}

type ProbesList []*Probe
