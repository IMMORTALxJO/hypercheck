package probe

import (
	"testing"

	"gotest.tools/assert"
)

func TestString(t *testing.T) {
	probe := NewString("teststring")
	assert.Equal(t, probe.GetType(), StringType)
	assert.Assert(t, GetProbeResult(probe, "", "", "==", "teststring"))
	assert.Assert(t, GetProbeResult(probe, "", "", "!=", "abasd"))
	assert.Assert(t, GetProbeResult(probe, "", "", "~", "tst"))
	assert.Assert(t, GetProbeResult(probe, "", "", "~=", "^teststring$"))
	assert.Assert(t, GetProbeResult(probe, "", "", "~=", "\\w+"))
	assert.Assert(t, GetProbeResult(probe, "", "", "~=", ".*test.*"))
	assert.Assert(t, GetProbeResult(probe, "", "", "~!", ".*abasd.*"))
	assert.Assert(t, GetProbeResult(probe, "", "length", "==", "10"))

	assert.Assert(t, !GetProbeResult(probe, "", "", "!=", "teststring"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "==", "abasd"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "~", "abasd"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "~=", "^abasd$"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "~=", ".*abasd.*"))
	assert.Assert(t, !GetProbeResult(probe, "", "", "~!", ".*test.*"))

	assert.Assert(t, GetProbeMsg(probe, "", "", "badop", "") != "")
	assert.Assert(t, GetProbeMsg(probe, "", "", "~=", `\`) != "")
	assert.Assert(t, GetProbeMsg(probe, "", "", "~!", `\`) != "")
}
