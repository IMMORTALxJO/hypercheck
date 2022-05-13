package dns

import "probe/probe"

func GenerateProbe(input string) (probe.Probe, string) {
	data := getDnsWrapper(input)
	dnsProbe := probe.NewMap()
	dnsProbe.Add("online", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(data.GetOnline())
	}))
	dnsProbe.Add("offline", probe.NewGenerator(func() probe.Probe {
		return probe.NewBool(!data.GetOnline())
	}))
	dnsProbe.Add("A", probe.NewGenerator(func() probe.Probe {
		records := probe.NewList()
		for _, ip := range data.GetA() {
			records.Add(probe.NewString(ip))
		}
		return records
	}))
	dnsProbe.Add("NS", probe.NewGenerator(func() probe.Probe {
		records := probe.NewList()
		for _, ip := range data.GetNS() {
			records.Add(probe.NewString(ip))
		}
		return records
	}))
	dnsProbe.Add("TXT", probe.NewGenerator(func() probe.Probe {
		records := probe.NewList()
		for _, ip := range data.GetTXT() {
			records.Add(probe.NewString(ip))
		}
		return records
	}))
	dnsProbe.Add("MX", probe.NewGenerator(func() probe.Probe {
		records := probe.NewList()
		for _, ip := range data.GetMX() {
			records.Add(probe.NewString(ip))
		}
		return records
	}))

	dnsProbe.Add("CNAME", probe.NewGenerator(func() probe.Probe {
		return probe.NewString(data.GetCNAME())
	}))

	return dnsProbe, ""
}
