package http

import (
	"hypercheck/probe/types"
	probe "hypercheck/probe/types"
)

const Name = "HTTP"

func GenerateProbe(input string) (types.Probe, string) {
	data := getHttpWrapper(input)
	httpProbe := types.NewMap("Check http resource")
	httpProbe.Add("code", types.NewGenerator("response status code", types.NumberType, func() types.Probe {
		return probe.NewNumber("", data.GetCode(), "int")
	}))
	httpProbe.Add("content", types.NewGenerator("response content", types.StringType, func() types.Probe {
		return probe.NewString("", data.GetContent())
	}))
	httpProbe.Add("online", types.NewGenerator("is online", types.BoolType, func() types.Probe {
		return probe.NewBool("", data.GetOnline())
	}))
	httpProbe.Add("offline", types.NewGenerator("is offline", types.BoolType, func() types.Probe {
		return probe.NewBool("", !data.GetOnline())
	}))
	httpProbe.Add("headers", types.NewGenerator("headers content", "List[String]", func() types.Probe {
		headersProbe := probe.NewList("")
		for _, h := range data.GetHeaders() {
			headersProbe.Add(probe.NewString("", h))
		}
		return headersProbe
	}))
	return httpProbe, ""
}
