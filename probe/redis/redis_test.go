package redis

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

func TestRedis(t *testing.T) {
	// listen redis port
	assert.Assert(t, getProbeResult(getProbe("localhost:6379"), "online", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("localhost:6379"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("localhost:8080"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("localhost:8080"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("localhost:8081"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("localhost:8081"), "online", "", "", ""))

}
