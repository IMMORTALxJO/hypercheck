package main

import (
	"os"
	"probe/cli"
	fsProbe "probe/probe/fs"
	httpProbe "probe/probe/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetLevel(log.DebugLevel)
	globalResult := true
	if len(os.Args) > 0 && (len(os.Args)-1)%3 != 0 {
		log.Errorf("Wrong number of attributes %d", len(os.Args))
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		cliArg := os.Args[i]
		log.Debugf("cliArg: %s", cliArg)
		probe, _ := httpProbe.GenerateProbe("")
		target := os.Args[i+2]
		err := ""
		switch cliArg {
		case "--fs":
			probe, err = fsProbe.GenerateProbe(target)
			if err != "" {
				log.Error(err)
				os.Exit(1)
			}
		case "--http":
			probe, err = httpProbe.GenerateProbe(target)
			if err != "" {
				log.Error(err)
				os.Exit(1)
			}
		default:
			log.Errorf("Unknown probe '%s'", cliArg)
			os.Exit(1)
		}
		for _, probeInput := range cli.ParseArguments(os.Args[i+1]) {
			result, msg := probe.Up(probeInput)
			if result {
				log.Infof("%s %s", target, probeInput.ToString())
			} else {
				log.Errorf("%s %s", target, msg)
				globalResult = false
			}
		}
		i += 2
	}
	if !globalResult {
		log.Error("probes failed")
		os.Exit(1)
	}
	log.Info("probes succeed")
}
