package main

import (
	"os"

	"goer-startup/internal/goerctl/cmd"
)

func main() {
	command := cmd.NewDefaultGoerCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
