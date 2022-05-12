package http

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestHttp(t *testing.T) {
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/status/200"), "code", "", "==", "200"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/basic-auth"), "code", "", "==", "401"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman:password@postman-echo.com/basic-auth"), "code", "", "==", "200"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/status/200"), "online", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/status/200"), "content", "", "==", "{\"status\":200}"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/status/200"), "content", "length", "==", "14"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://postman-echo.com/response-headers?foo1=bar1&foo2=bar2"), "headers", "any", "==", "Foo1: bar1"))

	assert.Assert(t, !probe.GetProbeResult(getProbe("https://offline-not-found.com"), "online", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://offline-not-found.com"), "offline", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://offline-not-found.com"), "code", "", "==", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://offline-not-found.com"), "content", "length", "==", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("https://offline-not-found.com"), "headers", "count", "==", "0"))
}
