package app

import (
	"log"

	"github.com/fatih/color"
)

func infof(format string, a ...interface{}) {
	log.Print(color.CyanString(format, a...))
}

func warnf(format string, a ...interface{}) {
	log.Print(color.YellowString(format, a...))
}
