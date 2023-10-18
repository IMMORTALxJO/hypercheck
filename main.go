package main

import (
	"fmt"
	"hypercheck/probe"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No command specified")
		os.Exit(1)
	}
	hypercheck := probe.New()
	queries := []string{}
	// Parse arguments
	// -q -- add next arg to queries
	// --tcp -- add next arg to hypercheck as tcp driverName

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--tcp" {
			log.Debugf("cli: Adding tcp driver %s", os.Args[i+1])
			hypercheck.Add("tcp", os.Args[i+1])
			i++
		} else if os.Args[i] == "--dns" {
			log.Debugf("cli: Adding dns driver %s", os.Args[i+1])
			hypercheck.Add("dns", os.Args[i+1])
			i++
		} else if os.Args[i] == "-v" {
			log.SetLevel(log.DebugLevel)
		} else {
			log.Debugf("cli: Adding query %s", os.Args[i])
			queries = append(queries, os.Args[i])
		}
	}
	hypercheck.Exec(queries)
	exitCode := hypercheck.Validate()
	os.Exit(exitCode)
}
