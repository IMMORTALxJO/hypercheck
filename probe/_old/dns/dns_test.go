package dns

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

func TestDns(t *testing.T) {
	assert.Assert(t, getProbeResult(getProbe("google.com"), "online", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("google.com"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("google.com"), "A", "count", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("gmail.com"), "MX", "count", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("google.com"), "NS", "count", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("mail.google.com"), "TXT", "count", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("research.swtch.com"), "CNAME", "length", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("research.swtch.com"), "CNAME", "", "==", "ghs.google.com."))

	// unknown domain
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("offline-not-found.com"), "online", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "A", "count", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "MX", "count", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "NS", "count", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "TXT", "count", "==", "0"))
	assert.Assert(t, getProbeResult(getProbe("offline-not-found.com"), "CNAME", "", "==", ""))
}
