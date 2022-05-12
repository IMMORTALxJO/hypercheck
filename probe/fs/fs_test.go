package fs

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func getProbe(pattern string) probe.Probe {
	probe, _ := GenerateProbe(pattern)
	return probe
}

func TestFS(t *testing.T) {
	fsProbe := getProbe("./assets/**")
	assert.Equal(t, fsProbe.GetType(), "Map")
	assert.Assert(t, probe.GetProbeResult(fsProbe, "count", "", ">", "0"))
	assert.Assert(t, probe.GetProbeResult(fsProbe, "count", "", "==", "7"))
	assert.Assert(t, probe.GetProbeResult(fsProbe, "exists", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/file_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/file_*"), "size", "sum", "==", "10Kb"))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/dir_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/dir_5kb"), "size", "", "==", "5Kb"))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/deep_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, probe.GetProbeResult(getProbe("./assets/d*"), "size", "sum", "==", "16Kb"))

	// check ANY condition
	assert.Assert(t, !probe.GetProbeResult(fsProbe, "dir", "", "", ""))
	assert.Assert(t, !probe.GetProbeResult(fsProbe, "regular", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(fsProbe, "dir", "any", "", ""))
	assert.Assert(t, probe.GetProbeResult(fsProbe, "regular", "any", "", ""))

}
