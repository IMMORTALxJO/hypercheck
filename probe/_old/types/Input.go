package types

type Input struct {
	Key        string
	Aggregator string
	Operator   string
	Value      string
}

func (i *Input) ToString() string {
	result := i.Key
	if i.Aggregator != "" {
		result += ":" + i.Aggregator
	}
	return result + i.Operator + i.Value
}

func NewProbeInput(key string, agg string, op string, target string) *Input {
	return &Input{Key: key, Aggregator: agg, Operator: op, Value: target}
}
