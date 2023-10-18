package tcp

import (
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestNewTcp(t *testing.T) {
	localhost := NewItem("localhost:1234")
	assert.DeepEqual(t, localhost, &Item{
		Address:     "localhost:1234",
		Message:     notCheckedMessage,
		LatencyNano: 0,
		Failed:      false,
		Checked:     false,
	})
	assert.Equal(t, localhost.GetMessage(), fmt.Sprintf("localhost:1234 - %s", notCheckedMessage))
	assert.Equal(t, localhost.IsFailed(), false)
	assert.Equal(t, localhost.TableName(), tableName)
}

func TestTcp(t *testing.T) {
	// mock tcp server
	go func() {
		ln, err := net.Listen("tcp", ":8000")
		if err != nil {
			panic(err)
		}
		defer ln.Close()
		for {
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			conn.Close()
		}
	}()

	offlineCheck := &Item{
		Address: "localhost:1234",
	}
	offlineCheck.Enrich()
	assert.Check(t, offlineCheck.Checked, "offline check should be marked as checked")
	assert.Check(t, offlineCheck.IsFailed(), "port 1234 should be closed")
	assert.Check(t, offlineCheck.GetMessage() != fmt.Sprintf("localhost:1234 - %s", checkIsUpMessage), "offline check should contain error message")
	assert.Check(t, offlineCheck.LatencyNano > 0, "offline check should not have latency 0")

	onlineCheck := &Item{
		Address: "localhost:8000",
	}
	onlineCheck.Enrich()
	assert.Check(t, offlineCheck.Checked, "online check should be marked as checked")
	assert.Check(t, !onlineCheck.IsFailed(), "port 8000 should be open")
	assert.DeepEqual(t, onlineCheck.GetMessage(), fmt.Sprintf("%s - %s", onlineCheck.Address, checkIsUpMessage))
	assert.Check(t, offlineCheck.LatencyNano > 0, "online check should not have latency 0")
}
