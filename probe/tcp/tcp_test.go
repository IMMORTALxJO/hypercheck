package tcp

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

func TestTcp(t *testing.T) {
	// listen ports
	assert.Assert(t, getProbeResult(getProbe("127.0.0.1:8080"), "online", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("127.0.0.1:8080"), "latency", "", ">=", "0"))
	assert.Assert(t, getProbeResult(getProbe("127.0.0.1:8080"), "latency", "", "<", "1h"))
	assert.Assert(t, !getProbeResult(getProbe("127.0.0.1:8080"), "offline", "", "", ""))

	// not listen ports
	assert.Assert(t, getProbeResult(getProbe("127.0.0.1:8081"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("127.0.0.1:8081"), "online", "", "", ""))

	// bad format ports
	assert.Assert(t, getProbeResult(getProbe("127.0.0.1:123456"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("127.0.0.1:123456"), "online", "", "", ""))

}
