package fs

import (
	"hypercheck/probe/types"
	"testing"

	log "github.com/sirupsen/logrus"

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

func TestFS(t *testing.T) {
	fsProbe := getProbe("./assets/**")
	assert.Equal(t, fsProbe.GetType(), "Map")
	log.SetLevel(log.DebugLevel)
	assert.Assert(t, getProbeResult(fsProbe, "count", "", ">", "0"))
	assert.Assert(t, getProbeResult(fsProbe, "count", "", "==", "7"))
	assert.Assert(t, getProbeResult(fsProbe, "exists", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("./assets/file_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/file_*"), "size", "sum", "==", "10Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/dir_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/dir_5kb"), "size", "", "==", "5Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/deep_1kb"), "size", "", "==", "1Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/d*"), "size", "sum", "==", "16Kb"))
	assert.Assert(t, getProbeResult(getProbe("./assets/deep_1kb"), "uid", "", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("./assets/deep_1kb"), "gid", "", ">", "0"))
	assert.Assert(t, getProbeResult(getProbe("./assets/deep_1kb"), "user", "", "!=", ""))
	assert.Assert(t, getProbeResult(getProbe("./assets/deep_1kb"), "group", "", "!=", ""))

	// check ANY condition
	assert.Assert(t, !getProbeResult(fsProbe, "dir", "", "", ""))
	assert.Assert(t, !getProbeResult(fsProbe, "regular", "", "", ""))
	assert.Assert(t, getProbeResult(fsProbe, "dir", "any", "", ""))
	assert.Assert(t, getProbeResult(fsProbe, "regular", "any", "", ""))

}
