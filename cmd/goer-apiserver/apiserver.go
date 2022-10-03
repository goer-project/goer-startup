package main

import (
	"math/rand"
	"time"

	apiserver "goer-startup/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	apiserver.NewApp("goer-apiserver").Run()
}
