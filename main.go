package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/skyscrapers/apt-package-resource/command"
	"github.com/smira/aptly/aptly"
	"github.com/smira/aptly/cmd"
)

// Version variable, filled in at link time
var Version string

func main() {
	if Version == "" {
		Version = "unknown"
	}

	aptly.Version = Version

	rand.Seed(time.Now().UnixNano())

	os.Exit(cmd.Run(command.RootCommand(), os.Args[1:], true))
}
