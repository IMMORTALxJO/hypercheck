package http

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getHttpWrapper(input)
	httpProbe := probe.NewMap()
	httpProbe.Add("code", probe.NewNumber(data.GetCode(), "int"))
	httpProbe.Add("content", probe.NewString(data.GetContent()))
	httpProbe.Add("online", probe.NewBool(data.GetOnline()))
	httpProbe.Add("offline", probe.NewBool(!data.GetOnline()))
	headersProbe := probe.NewList()
	for _, h := range data.GetHeaders() {
		headersProbe.Add(probe.NewString(h))
	}
	httpProbe.Add("headers", headersProbe)
	return httpProbe, ""
}
