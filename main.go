package main

import (
	"probe/probes/fs"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fs.Probe("./probes/fs/assets/dir_10kb*", "size==10Kb")
}
