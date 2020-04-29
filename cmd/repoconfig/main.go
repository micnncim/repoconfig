package main

import (
	"fmt"
	"os"

	"github.com/micnncim/repoconfig/pkg/app"
)

func main() {
	if err := app.NewCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute command: %s", err)
		os.Exit(1)
	}
}
