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
	assert.Assert(t, probe.GetProbeResult(getProbe("8.8.8.8:53"), "online", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("8.8.8.8:77777"), "offline", "", "", ""))
}
