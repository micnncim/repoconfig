package app

import (
	"log"

	"github.com/fatih/color"
)

func infof(format string, a ...interface{}) {
	log.Printf(color.CyanString(format, a...))
}
