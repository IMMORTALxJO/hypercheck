package auto

import (
	"probe/probe"
	dbProbe "probe/probe/db"
	dnsProbe "probe/probe/dns"
	httpProbe "probe/probe/http"
	redisProbe "probe/probe/redis"
	tcpProbe "probe/probe/tcp"

	schemeDetector "github.com/IMMORTALxJO/scheme-detector"
)

func GenerateProbe(input string) (probe.Probe, string) {
	autoProbe := probe.NewList()
	for _, scheme := range schemeDetector.FromEnv() {
		if scheme.IsDNSName() {
			dns, _ := dnsProbe.GenerateProbe(scheme.URL.Hostname())
			autoProbe.Add(probe.NewParametrized(dns, probe.NewProbeInput("online", "", "", "")))
		}

		if scheme.URL.Scheme == "http" || scheme.URL.Scheme == "https" {
			http, _ := httpProbe.GenerateProbe(scheme.URL.String())
			autoProbe.Add(probe.NewParametrized(http, probe.NewProbeInput("online", "", "", "")))
			autoProbe.Add(probe.NewParametrized(http, probe.NewProbeInput("online", "", "", "")))
			autoProbe.Add(probe.NewParametrized(http, probe.NewProbeInput("code", "", ">=", "200")))
			autoProbe.Add(probe.NewParametrized(http, probe.NewProbeInput("code", "", "<", "500")))
		}
		if scheme.URL.Scheme == "redis" {
			tcp, _ := tcpProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(probe.NewParametrized(tcp, probe.NewProbeInput("online", "", "", "")))
			redis, _ := redisProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(probe.NewParametrized(redis, probe.NewProbeInput("online", "", "", "")))
		}
		if scheme.URL.Scheme == "postgres" || scheme.URL.Scheme == "mysql" {
			tcp, _ := tcpProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(probe.NewParametrized(tcp, probe.NewProbeInput("online", "", "", "")))
			db, _ := dbProbe.GenerateProbe(scheme.URL.String())
			autoProbe.Add(probe.NewParametrized(db, probe.NewProbeInput("online", "", "", "")))
		}

	}
	return autoProbe, ""
}
