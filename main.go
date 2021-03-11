package main

import (
	"stockbit/cnf/env"
	"stockbit/handle"
)

func main() {
	env.LoadEnv()

	// Init dependencies
	serv := handle.MakeHandler()

	// start echo server
	serv.StartServer()

	// Shutdown with gracefull handler
	serv.ShutdownServer()
}
