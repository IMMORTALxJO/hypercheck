package main

import (
	"fmt"
	"hypercheck/probe"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	for os.Args[0] == "hypercheck" {
		os.Args = os.Args[1:]
	}
	if len(os.Args) == 0 {
		fmt.Println("No command specified")
		os.Exit(1)
	}
	hypercheck := probe.New()
	for _, arg := range os.Args[1:] {
		if arg == "-v" {
			log.SetLevel(log.DebugLevel)
		} else {
			hypercheck.Add("tcp", arg)
		}
	}
	hypercheck.Run()
	hypercheck.Validate()
}
