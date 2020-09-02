package main

import (
	"fmt"
	"github.com/TheComputerDan/sentinel_server/internal/api"
	"github.com/TheComputerDan/sentinel_server/internal/config"
	"github.com/TheComputerDan/sentinel_server/internal/data"
	"github.com/TheComputerDan/sentinel_server/internal/host"
	flag "github.com/spf13/pflag"
)

func main() {
	//svr.Start()

	confSetup := flag.BoolP(
		"setup",
		"s",
		false,
		"Run config setup.",
		)

	runApi := flag.BoolP(
		"api",
		"a",
		false,
		"Starts the server REST API",
		)

	dbTest := flag.BoolP(
		"dbt",
		"t",
		false,
		"Test DB function",
		)

	hostTest := flag.BoolP(
		"ht",
		"o",
		false,
		"Test Host Info",
		)

	flag.Parse()

	if *confSetup == true {
		config.Generate()
	}

	if *runApi == true{
		api.Start()
	}

	if *dbTest == true{
		data.Test()
	}

	if *hostTest == true{
		hostInfo := host.Info{}
		hostInfo.Init()
		fmt.Println(hostInfo)
	}

}
