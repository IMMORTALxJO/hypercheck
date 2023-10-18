package dns

import (
	"fmt"
	"net"
	"testing"

	"github.com/foxcpp/go-mockdns"
	"gotest.tools/assert"
)

func TestNewDns(t *testing.T) {
	record := NewItem("google.com:TXT")
	assert.DeepEqual(t, record, &Item{
		Address: "google.com",
		RType:   "TXT",
		Message: notCheckedMessage,
		Failed:  false,
		Checked: false,
		Value:   "",
	})
	assert.Equal(t, record.GetMessage(), fmt.Sprintf("google.com:TXT - %s", notCheckedMessage))
	assert.Equal(t, record.IsFailed(), false)
	assert.Equal(t, record.TableName(), tableName)

	record = NewItem("google.com")
	assert.Equal(t, record.GetMessage(), fmt.Sprintf("google.com:A - %s", notCheckedMessage), "input with no rType should fallback to A record")
}

func TestDnsOnline(t *testing.T) {
	domain := "it-is-online.com"
	srv, _ := mockdns.NewServer(map[string]mockdns.Zone{
		"it-is-online.com.": {
			A:     []string{"1.2.3.4", "2.3.4.5"},
			TXT:   []string{"custom txt record"},
			CNAME: "cname.it-is-online.com.",
		},
	}, false)
	defer srv.Close()

	srv.PatchNet(net.DefaultResolver)
	defer mockdns.UnpatchNet(net.DefaultResolver)

	record := &Item{
		Address: domain,
		RType:   "A",
	}
	record.Enrich()
	assert.DeepEqual(t, record.GetMessage(), fmt.Sprintf("%s:%s - %s", record.Address, record.RType, checkIsUpMessage))
	assert.Check(t, record.Checked, "it-is-online.com check should be marked as checked")
	assert.Equal(t, record.Value, "1.2.3.4\n2.3.4.5", "it-is-online.com should have A record value")
	assert.Check(t, !record.IsFailed(), "it-is-online.com should be exists")
}

func TestDnsOffline(t *testing.T) {
	offlineCheck := &Item{
		Address: "it-is-offline.com",
		RType:   "A",
	}
	offlineCheck.Enrich()
	assert.DeepEqual(t, offlineCheck.GetMessage(), fmt.Sprintf("%s:%s - %s", offlineCheck.Address, offlineCheck.RType, offlineCheck.Message))
	assert.Check(t, offlineCheck.Checked, "offline check should be marked as checked")
	assert.Check(t, offlineCheck.Value == "", "offline check should not have value")
	assert.Check(t, offlineCheck.IsFailed(), "offline check should be failed")
}
