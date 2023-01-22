package main

import (
	"os"

	"goer-startup/internal/apiserver"
)

func main() {
	command := apiserver.NewAppCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
