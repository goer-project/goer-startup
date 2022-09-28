package main

import (
	"math/rand"
	"time"

	apiserver "goer-startup/internal/goer-apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	apiserver.NewApp("goer-apiserver").Run()
}
