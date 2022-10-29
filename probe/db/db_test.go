package db

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

func TestPostgres(t *testing.T) {
	// listen port
	assert.Assert(t, getProbeResult(getProbe("postgres://user:password@localhost:5432/postgres?sslmode=disable"), "online", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("postgres://user:password@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("postgres://baduser:password@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("postgres://user:badpassword@localhost:5432/postgres?sslmode=disable"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("postgres://user:password@localhost:8080/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("postgres://user:password@localhost:8080/postgres?sslmode=disable"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("postgres://user:password@localhost:8081/postgres?sslmode=disable"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("postgres://user:password@localhost:8081/postgres?sslmode=disable"), "online", "", "", ""))

}

func TestMysql(t *testing.T) {
	// listen port
	assert.Assert(t, getProbeResult(getProbe("mysql://user:password@localhost:3306/database"), "online", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("mysql://user:password@localhost:3306/database"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("mysql://baduser:password@localhost:3306/database"), "offline", "", "", ""))
	assert.Assert(t, getProbeResult(getProbe("mysql://user:badpassword@localhost:3306/database"), "offline", "", "", ""))

	// listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("mysql://user:password@localhost:8080/database"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("mysql://user:password@localhost:8080/database"), "online", "", "", ""))

	// not listen port but not redis
	assert.Assert(t, getProbeResult(getProbe("mysql://user:password@localhost:8081/database"), "offline", "", "", ""))
	assert.Assert(t, !getProbeResult(getProbe("mysql://user:password@localhost:8081/database"), "online", "", "", ""))

}
