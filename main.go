package main

import (
	"probe/probe/fs"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fs.Probe("./probe/fs/assets/dir_10kb*", "size==10Kb")
}
