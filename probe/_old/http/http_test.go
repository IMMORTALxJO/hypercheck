package http

import (
	"hypercheck/probe/types"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) types.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func getProbeResult(probe types.Probe, key string, agg string, op string, target string) bool {
	result, _ := probe.Up(types.NewProbeInput(key, agg, op, target))
	return result
}

func getProbeMsg(probe types.Probe, key string, agg string, op string, target string) string {
	_, msg := probe.Up(types.NewProbeInput(key, agg, op, target))
	return msg
}

func TestHttp(t *testing.T) {
	localNginx := "http://localhost:8080"
	assert.Assert(t, getProbeResult(getProbe(localNginx), "code", "", "==", "200"))
	assert.Assert(t, getProbeResult(getProbe(localNginx), "online", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe(localNginx), "content", "", "~", "Welcome to nginx!"))
	assert.Assert(t, getProbeResult(getProbe(localNginx), "content", "length", "==", "612"))

	assert.Assert(t, getProbeResult(getProbe("https://postman-echo.com/basic-auth"), "code", "", "==", "401"))
	assert.Assert(t, getProbeResult(getProbe("https://postman:password@postman-echo.com/basic-auth"), "code", "", "==", "200"))
	assert.Assert(t, getProbeResult(getProbe("https://postman-echo.com/response-headers?foo1=bar1&foo2=bar2"), "headers", "any", "==", "Foo1: bar1"))
	assert.Assert(t, !getProbeResult(getProbe("https://offline-not-found.com"), "online", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("https://offline-not-found.com"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("https://offline-not-found.com"), "code", "", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("https://offline-not-found.com"), "content", "length", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("https://offline-not-found.com"), "headers", "count", "==", "0"))
}

func BenchmarkHTTP(b *testing.B) {
	localNginx := "http://localhost:8080"
	for i := 0; i < b.N; i++ {
		getProbeResult(getProbe(localNginx), "code", "", "==", "200")
	}
}
