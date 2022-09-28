package main

import (
	"goer-startup/internal/goerctl/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultGoerCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
