package probe

type Probe interface {
	Up(*Input) (bool, string)
	GetType() string
}

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

func GetProbeResult(probe Probe, key string, agg string, op string, target string) bool {
	result, _ := probe.Up(&Input{Key: key, Aggregator: agg, Operator: op, Value: target})
	return result
}

func GetProbeMsg(probe Probe, key string, agg string, op string, target string) string {
	_, msg := probe.Up(&Input{Key: key, Aggregator: agg, Operator: op, Value: target})
	return msg
}
