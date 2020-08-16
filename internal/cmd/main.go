package main

import (
	// "fmt"
	// client "github.com/TheComputerDan/heimdall/docker_agent/docker/comms"
	svr "github.com/TheComputerDan/heimdall_server/internal/docker/comms"
	_ "github.com/TheComputerDan/heimdall_server/internal/host"
	// app "github.com/TheComputerDan/heimdall/docker_agent/web/api"
)

func main() {
	// app.Start()
	svr.Start()
	// client.Start()

}
