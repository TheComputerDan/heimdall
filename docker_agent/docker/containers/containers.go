package containers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Instantiate connects to the local docker socket and returns `[]type.Container` from the official `docker/api/types` package
func Instantiate() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	return containers
}

// Info Returns JSON with the
func Info(containers []types.Container) {
	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.Names[0])
		fmt.Println(container.Ports)
		fmt.Println(container.Status)
		fmt.Println(container.State)
	}
}
