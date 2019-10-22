package main

import (
	config "./config"
	Routes "./router"
)

var err error

func main() {

	r := Routes.SetupRouter()
	// running
	r.Run(config.PORT)
}
