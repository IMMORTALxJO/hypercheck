package tcp

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestTcp(t *testing.T) {
	// listen ports
	assert.Assert(t, probe.GetProbeResult(getProbe("127.0.0.1:8080"), "online", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("127.0.0.1:8080"), "latency", "", ">=", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("127.0.0.1:8080"), "latency", "", "<", "1h"))
	assert.Assert(t, !probe.GetProbeResult(getProbe("127.0.0.1:8080"), "offline", "", "", ""))

	// not listen ports
	assert.Assert(t, probe.GetProbeResult(getProbe("127.0.0.1:8081"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("127.0.0.1:8081"), "online", "", "", ""))

	// bad format ports
	assert.Assert(t, probe.GetProbeResult(getProbe("127.0.0.1:123456"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("127.0.0.1:123456"), "online", "", "", ""))

}
