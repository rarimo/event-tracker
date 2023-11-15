package main

import (
	"os"

	"github.com/rarimo/event-tracker/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
