package http

import (
	"fmt"
	"hypercheck/probe/types"
	probe "hypercheck/probe/types"
)

const Name = "HTTP"

func GenerateProbe(input string) (types.Probe, string) {
	data := getHttpWrapper(input)
	description := "Check HTTP response"
	if len(input) > 0 {
		description = fmt.Sprintf("HTTP %s", input)
	}

	httpProbe := types.NewMap(description, len(input) > 0)
	httpProbe.Add("code", types.NewGenerator("response status code", types.NumberType, func() types.Probe {
		return probe.NewNumber("", data.getCode(), "int")
	}))
	httpProbe.Add("content", types.NewGenerator("response content", types.StringType, func() types.Probe {
		return probe.NewString("", data.getContent())
	}))
	httpProbe.Add("online", types.NewGenerator("status code 200", types.BoolType, func() types.Probe {
		return probe.NewBool("", data.getOnline())
	}))
	httpProbe.Add("offline", types.NewGenerator("status code is not 200", types.BoolType, func() types.Probe {
		return probe.NewBool("", !data.getOnline())
	}))
	httpProbe.Add("headers", types.NewGenerator("headers content", "List[String]", func() types.Probe {
		headersProbe := probe.NewList("")
		for _, h := range data.getHeaders() {
			headersProbe.Add(probe.NewString("", h))
		}
		return headersProbe
	}))
	return httpProbe, ""
}
