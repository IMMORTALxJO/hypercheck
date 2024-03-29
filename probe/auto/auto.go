package auto

import (
	dbProbe "hypercheck/probe/db"
	dnsProbe "hypercheck/probe/dns"
	httpProbe "hypercheck/probe/http"
	redisProbe "hypercheck/probe/redis"
	tcpProbe "hypercheck/probe/tcp"
	"hypercheck/probe/types"
	"os"

	log "github.com/sirupsen/logrus"

	schemeDetector "github.com/IMMORTALxJO/scheme-detector"
)

const Name = "Auto"

func GenerateProbe() (types.Probe, string) {
	os.Setenv("SCHEME_DETECTOR_EXCLUDE", os.Getenv("HYPERCHECK_ENV_EXCLUDE"))
	autoProbe := types.NewList("Generate probes automaticaly based on current environment variables")
	for _, scheme := range schemeDetector.FromEnv() {
		if scheme.IsDNSName() {
			dns, _ := dnsProbe.GenerateProbe(scheme.URL.Hostname())
			autoProbe.Add(types.NewParametrized(dns, types.NewProbeInput("A", "count", ">", "0")))
		}

		if scheme.URL.Scheme == "http" || scheme.URL.Scheme == "https" {
			log.Debugf("Found HTTP: %s", scheme.URL.String())
			http, _ := httpProbe.GenerateProbe(scheme.URL.String())
			autoProbe.Add(types.NewParametrized(http, types.NewProbeInput("code", "", "==", "200")))
		}
		if scheme.URL.Scheme == "redis" {
			log.Debugf("Found Redis: %s", scheme.URL.String())
			tcp, _ := tcpProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(types.NewParametrized(tcp, types.NewProbeInput("online", "", "", "")))
			redis, _ := redisProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(types.NewParametrized(redis, types.NewProbeInput("online", "", "", "")))
		}
		if scheme.URL.Scheme == "postgres" || scheme.URL.Scheme == "mysql" {
			log.Debugf("Found DB: %s", scheme.String())
			tcp, _ := tcpProbe.GenerateProbe(scheme.URL.Hostname() + ":" + scheme.URL.Port())
			autoProbe.Add(types.NewParametrized(tcp, types.NewProbeInput("online", "", "", "")))
			db, _ := dbProbe.GenerateProbe(scheme.String())
			autoProbe.Add(types.NewParametrized(db, types.NewProbeInput("online", "", "", "")))
		}

	}
	return autoProbe, ""
}
