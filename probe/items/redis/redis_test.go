package redis

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNewRedis(t *testing.T) {
	record := NewItem("localhost")
	assert.DeepEqual(t, record, &Item{
		Address: "localhost:6379",
		Message: notCheckedMessage,
		Failed:  false,
		Checked: false,
	})
	assert.Equal(t, record.GetMessage(), fmt.Sprintf("localhost:6379 - %s", notCheckedMessage))
	assert.Equal(t, record.IsFailed(), false)
	assert.Equal(t, record.TableName(), tableName)

}

func TestRedisOnline(t *testing.T) {
	record := &Item{
		Address: "localhost:6379",
	}
	record.Enrich()
	assert.DeepEqual(t, record.GetMessage(), fmt.Sprintf("%s - %s", record.Address, checkIsUpMessage))
	assert.Check(t, record.Checked, "localhost:6379 check should be marked as checked")
	assert.Check(t, record.Failed == false, "localhost:6379 should be online")
}

func TestRedisOffline(t *testing.T) {
	record := &Item{
		Address: "localhost:8080",
	}
	record.Enrich()
	assert.Check(t, record.Checked, "localhost:8080 check should be marked as checked")
	assert.Check(t, record.Failed, "localhost:8080 should be failed")
}
