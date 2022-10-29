package dns

import "hypercheck/probe/types"

const Name = "DNS"

func GenerateProbe(input string) (types.Probe, string) {
	data := getDnsWrapper(input)
	dnsProbe := types.NewMap("Check dns query response")
	dnsProbe.Add("online", types.NewGenerator("A record is not empty", types.BoolType, func() types.Probe {
		return types.NewBool("", data.GetOnline())
	}))
	dnsProbe.Add("offline", types.NewGenerator("A record is empty", types.BoolType, func() types.Probe {
		return types.NewBool("", !data.GetOnline())
	}))
	dnsProbe.Add("A", types.NewGenerator("A record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.GetA() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("NS", types.NewGenerator("NS record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.GetNS() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("TXT", types.NewGenerator("TXT record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.GetTXT() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))
	dnsProbe.Add("MX", types.NewGenerator("MX record content", "List[string]", func() types.Probe {
		records := types.NewList("")
		for _, ip := range data.GetMX() {
			records.Add(types.NewString("", ip))
		}
		return records
	}))

	dnsProbe.Add("CNAME", types.NewGenerator("CNAME record content", types.StringType, func() types.Probe {
		return types.NewString("", data.GetCNAME())
	}))

	return dnsProbe, ""
}
