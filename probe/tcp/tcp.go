package tcp

import "hypercheck/probe/types"

const Name = "TCP"

func GenerateProbe(input string) (types.Probe, string) {
	data := getTcpWrapper(input)
	tcpProbe := types.NewMap("Check tcp port")
	tcpProbe.Add("online", types.NewGenerator("is reachible", types.BoolType, func() types.Probe {
		return types.NewBool("", data.GetOnline())
	}))
	tcpProbe.Add("offline", types.NewGenerator("is unreachible", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.GetOnline())
	}))
	tcpProbe.Add("latency", types.NewGenerator("duration", types.NumberType, func() types.Probe {
		return types.NewNumber("", data.GetLatency(), "duration")
	}))
	return tcpProbe, ""
}
