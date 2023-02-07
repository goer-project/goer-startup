package main

import (
	"os"

	"goer-startup/internal/watcher"
)

func main() {
	command := watcher.NewWatcherCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
