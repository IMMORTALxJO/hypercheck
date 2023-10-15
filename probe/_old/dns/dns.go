package dns

import (
	"fmt"
	"hypercheck/probe/types"
)

const Name = "DNS"

func GenerateProbe(input string) (types.Probe, string) {
	data := getDnsWrapper(input)
	description := "Check DNS query response"
	if len(input) > 0 {
		description = fmt.Sprintf("DNS %s", input)
	}
	dnsProbe := types.NewMap(description, len(input) > 0)
	dnsProbe.Add("online", types.NewGenerator("A record is not empty", types.BoolType, func() types.Probe {
		return types.NewBool("", data.getOnline())
	}))
	dnsProbe.Add("offline", types.NewGenerator("A record is empty", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.getOnline())
	}))
	dnsProbe.Add("A", types.NewGenerator("A record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.getA() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("NS", types.NewGenerator("NS record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.getNS() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("TXT", types.NewGenerator("TXT record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.getTXT() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("MX", types.NewGenerator("MX record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.getMX() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))

	dnsProbe.Add("CNAME", types.NewGenerator("CNAME record content", types.StringType, func() types.Probe {
		return types.NewString("", data.getCNAME())
	}))

	return dnsProbe, ""
}
