package tcp

import (
	"fmt"
	"hypercheck/probe/types"
)

const Name = "TCP"

func GenerateProbe(input string) (types.Probe, string) {
	data := getTcpWrapper(input)
	description := "Check TCP address and port"
	if len(input) > 0 {
		description = fmt.Sprintf("TCP %s", input)
	}
	tcpProbe := types.NewMap(description, len(input) > 0)
	tcpProbe.Add("online", types.NewGenerator("is reachable", types.BoolType, func() types.Probe {
		return types.NewBool("", data.getOnline())
	}))
	tcpProbe.Add("offline", types.NewGenerator("is unreachable", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.getOnline())
	}))
	tcpProbe.Add("latency", types.NewGenerator("duration", types.NumberType, func() types.Probe {
		return types.NewNumber("", data.getLatency(), "duration")
	}))
	return tcpProbe, ""
}
