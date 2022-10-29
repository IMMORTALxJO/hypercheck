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
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	globalResult := true
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-v" {
			log.SetLevel(log.DebugLevel)
		}
		if os.Args[i] == "--help" {
			printHelp()
			os.Exit(0)
		}
	}
	i := 1
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
		probe, _ := httpProbe.GenerateProbe("")

		err := ""
		args := ""
		target := ""
		probeName := ""

		if len(os.Args) > i && !isFlag(os.Args[i]) {
			args = os.Args[i]
			log.Debugf("args: %s", args)
			i++
			log.Debugf("i := %d", i)
		}
		if len(os.Args) > i && !isFlag(os.Args[i]) {
			target = os.Args[i]
			log.Debugf("target: %s", args)
			i++
			log.Debugf("i := %d", i)
		}
		switch probeFlag {
		case "-v":
			continue
		case "--fs":
			probeName = fsProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "exists"
			}
			probe, err = fsProbe.GenerateProbe(target)
		case "--http":
			probeName = httpProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "online"
			}
			probe, err = httpProbe.GenerateProbe(target)
		case "--tcp":
			probeName = tcpProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "online"
			}
			probe, err = tcpProbe.GenerateProbe(target)
		case "--dns":
			probeName = dnsProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "online"
			}
			probe, err = dnsProbe.GenerateProbe(target)
		case "--redis":
			probeName = redisProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "online"
			}
			probe, err = redisProbe.GenerateProbe(target)
		case "--db":
			probeName = dbProbe.Name
			if target == "" && args != "" {
				target = args
				log.Debugf("probeTarget := probeArgs")
				args = "online"
			}
			probe, err = dbProbe.GenerateProbe(target)
		case "--auto":
			probeName = autoProbe.Name
			probe, err = autoProbe.GenerateProbe()
		}
		if err != "" {
			log.Error(err)
			os.Exit(1)
		}
		fmt.Printf("Checking '%s' %s ...\n", args, target)
		for _, probeInput := range cli.ParseArguments(args) {
			result, msg := probe.Up(probeInput)
			if result {
				fmt.Printf("\t\u2705  %s %s %s\n", probeName, probeInput.ToString(), target)
			} else {
				fmt.Printf("\t\u274C  %s %s %s ( %s )\n", probeName, probeInput.ToString(), target, msg)
				globalResult = false
			}
		}
	}
	if !globalResult {
		os.Exit(1)
	}
}

func isFlag(arg string) bool {
	switch arg {
	case "-v":
		return true
	case "--help":
		return true
	case "--fs":
		return true
	case "--http":
		return true
	case "--tcp":
		return true
	case "--dns":
		return true
	case "--redis":
		return true
	case "--db":
		return true
	case "--auto":
		return true
	}
	return false
}

func printHelp() {
	probe, _ := fsProbe.GenerateProbe("")
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
	probe, _ = autoProbe.GenerateProbe()
	fmt.Println("--auto\n", probe.GetDescription())
}
