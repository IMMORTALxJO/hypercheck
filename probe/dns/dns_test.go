package dns

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestDns(t *testing.T) {
	assert.Assert(t, probe.GetProbeResult(getProbe("google.com"), "A", "count", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("gmail.com"), "MX", "count", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("google.com"), "NS", "count", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("mail.google.com"), "TXT", "count", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("research.swtch.com"), "CNAME", "length", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(getProbe("research.swtch.com"), "CNAME", "", "==", "ghs.google.com."))
}
