package main

import (
	// "fmt"
	_ "github.com/TheComputerDan/heimdall/docker_agent/host"
	app "github.com/TheComputerDan/heimdall/docker_agent/web/api"
)

func main() {
	app.Start()
}
