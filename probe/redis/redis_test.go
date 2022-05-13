package redis

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestRedis(t *testing.T) {
	// listen redis port
	assert.Assert(t, probe.GetProbeResult(getProbe("localhost:6379"), "online", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("localhost:6379"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("localhost:8080"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("localhost:8080"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, probe.GetProbeResult(getProbe("localhost:8081"), "offline", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(getProbe("localhost:8081"), "online", "", "", ""))

}
