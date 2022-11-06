package main

import (
	"fmt"
	"hypercheck/cli"
	autoProbe "hypercheck/probe/auto"
	dbProbe "hypercheck/probe/db"
	dnsProbe "hypercheck/probe/dns"
	fsProbe "hypercheck/probe/fs"
	httpProbe "hypercheck/probe/http"
	redisProbe "hypercheck/probe/redis"
	tcpProbe "hypercheck/probe/tcp"
	types "hypercheck/probe/types"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	globalResult := true
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-v" {
			log.SetLevel(log.DebugLevel)
		}
		if os.Args[i] == "--help" || os.Args[i] == "-h" {
			printHelp()
			os.Exit(0)
		}
	}
	i := 1
	var globalProbes []types.Probe
	for i < len(os.Args) {
		log.Debugf("parseArg[%d]: %s", i, os.Args[i])

		probeFlag := os.Args[i]
		if !isFlag(probeFlag) {
			log.Errorf("Unknown argument '%s'", probeFlag)
			os.Exit(1)
		}
		log.Debugf("probeFlag: %s", probeFlag)
		i++
		log.Debugf("i := %d", i)

		var probe types.Probe
		args := ""
		target := ""
		err := ""
		for _, n := range []string{"args", "target"} {
			if len(os.Args) > i && !isFlag(os.Args[i]) {
				if args == "" {
					args = os.Args[i]
				} else {
					target = os.Args[i]
				}
				log.Debugf("%s: %s", n, args)
				i++
				log.Debugf("i := %d", i)
			}
		}
		if target == "" && args != "" {
			target = args
			log.Debugf("probeTarget := probeArgs")
			args = "online"
		}

		switch probeFlag {
		case "-v":
			continue
		case "--fs":
			if args == "online" {
				args = "exists"
			}
			probe, err = fsProbe.GenerateProbe(target)
		case "--http":
			probe, err = httpProbe.GenerateProbe(target)
		case "--tcp":
			probe, err = tcpProbe.GenerateProbe(target)
		case "--dns":
			probe, err = dnsProbe.GenerateProbe(target)
		case "--redis":
			probe, err = redisProbe.GenerateProbe(target)
		case "--db":
			probe, err = dbProbe.GenerateProbe(target)
		case "--auto":
			os.Setenv("SCHEME_DETECTOR_EXCLUDE", os.Getenv("HYPERCHECK_ENV_EXCLUDE"))
			probe, err = autoProbe.GenerateProbe()
		default:
			log.Errorf("Unknown argument '%s'", probeFlag)
			os.Exit(1)
		}
		if probe.GetType() == types.ListType {
			globalProbes = append(globalProbes, probe.(*types.List).GetValue()...)
		} else {
			for _, probeInput := range cli.ParseArguments(args) {
				globalProbes = append(globalProbes, types.NewParametrized(probe, probeInput))
			}
		}
		if err != "" {
			log.Error(err)
			os.Exit(1)
		}
	}
	for _, probe := range globalProbes {
		result, msg := probe.Up(types.NewProbeInput("", "", "", ""))
		if result {
			fmt.Printf("\u2705 %s %s\n", probe.GetDescription(), msg)
		} else {
			fmt.Printf("\u274C %s %s\n", probe.GetDescription(), msg)
			globalResult = false
		}
	}
	if !globalResult {
		os.Exit(1)
	}
}

func isFlag(arg string) bool {
	return arg[0] == '-'
}

func printHelp() {
	probe, _ := autoProbe.GenerateProbe()
	fmt.Println("--auto\n", probe.GetDescription())
	probe, _ = fsProbe.GenerateProbe("")
	fmt.Println("--fs\n", probe.GetDescription())
	probe, _ = httpProbe.GenerateProbe("")
	fmt.Println("--http\n", probe.GetDescription())
	probe, _ = tcpProbe.GenerateProbe("")
	fmt.Println("--tcp\n", probe.GetDescription())
	probe, _ = dnsProbe.GenerateProbe("")
	fmt.Println("--dns\n", probe.GetDescription())
	probe, _ = tcpProbe.GenerateProbe("")
	fmt.Println("--tcp\n", probe.GetDescription())
	probe, _ = redisProbe.GenerateProbe("")
	fmt.Println("--redis\n", probe.GetDescription())
	probe, _ = dbProbe.GenerateProbe("")
	fmt.Println("--db\n", probe.GetDescription())
}
