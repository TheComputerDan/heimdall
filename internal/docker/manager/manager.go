package manager

import (
	_ "fmt"
	"github.com/TheComputerDan/sentinel_server/internal/docker/connect"
	"github.com/TheComputerDan/sentinel_server/internal/host"
	"github.com/docker/docker/api/types"
)

type dockerHost struct {
	hostname string
	containers []types.Container
	images []types.ImageSummary
}

//loadLocal populates the dockerHost from the server alone.
func loadLocal() dockerHost {

	hostInfo:=host.Info{}
	hostInfo.Init()

	serverHostname := hostInfo.Hostname
	dockerContainers := connect.Containers()
	dockerImages := connect.Images()

	return dockerHost{hostname: serverHostname, containers: dockerContainers, images: dockerImages}
}

//loadLocal populates a slice of dockerHost to return and provide to the REST API.
func loadAll() []dockerHost {
	var dockerHosts []dockerHost

	dockerHosts = append(dockerHosts, loadLocal())

	//TODO Get from all hosts

	return dockerHosts
}