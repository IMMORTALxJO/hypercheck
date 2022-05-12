package http

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getHttpWrapper(input)
	httpProbe := probe.NewMap()
	httpProbe.Add("code", probe.NewGenerator(func() probe.Probe {
		return probe.NewNumber(data.GetCode(), "int")
	}))
	httpProbe.Add("content", probe.NewGenerator(func() probe.Probe {
		return probe.NewString(data.GetContent())
	}))
	httpProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetOnline())
	}))
	httpProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetOnline())
	}))
	httpProbe.Add("headers", probe.NewGenerator(func() probe.Probe {
		headersProbe := probe.NewList()
		for _, h := range data.GetHeaders() {
			headersProbe.Add(probe.NewString(h))
		}
		return headersProbe
	}))
	return httpProbe, ""
}
