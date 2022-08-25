package probe

type Probe interface {
	Up(*Input) (bool, string)
	GetType() string
}

func GetProbeResult(probe Probe, key string, agg string, op string, target string) bool {
	result, _ := probe.Up(NewProbeInput(key, agg, op, target))
	return result
}

func GetProbeMsg(probe Probe, key string, agg string, op string, target string) string {
	_, msg := probe.Up(NewProbeInput(key, agg, op, target))
	return msg
}
