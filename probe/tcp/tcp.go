package tcp

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getTcpWrapper(input)
	tcpProbe := probe.NewMap()
	tcpProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetOnline())
	}))
	tcpProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetOnline())
	}))
	tcpProbe.Add("latency", probe.NewGenerator(func() probe.Probe {
		return probe.NewNumber(data.GetLatency(), "duration")
	}))
	return tcpProbe, ""
}
