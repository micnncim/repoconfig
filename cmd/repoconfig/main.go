package main

import (
	"log"

	"github.com/micnncim/repoconfig/pkg/app"
)

func main() {
	log.SetFlags(0)

	cmd, err := app.NewCommand()
	if err != nil {
		log.Fatalf("failed to set up command: %s", err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %s", err)
	}
}
