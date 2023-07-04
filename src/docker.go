package src

import (
	"context"
	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"strings"
)

func ShortenContainerId(id string) string {
	return id[0:12]
}

func SanitizeContainer(containers []types.Container) []types.Container {

	for i := 0; i < len(containers); i++ {
		//replace backslash in container names
		for j := 0; j < len(containers[i].Names); j++ {
			containers[i].Names[j] = strings.Replace(containers[i].Names[j], "/", "", 1)
		}
	}

	return containers
}

func ContainerData() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		log.Fatal(err)
	}
	options := types.ContainerListOptions{
		All: true, // Include stopped containers as well
	}

	containers, err := cli.ContainerList(context.Background(), options)

	if err != nil {
		log.Fatal(err)
	}

	//trim container id before returning
	return SanitizeContainer(containers)
}

func StartContainer(containerId string) {
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Fatal(err)
	}

	err = cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})

	if err != nil {
		log.Fatal(err)
	}
}

func StopContainer(containerId string) {
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Fatal(err)
	}

	// options := container.StopOptions{}
	// err = cli.ContainerStop(context.Background(), containerId, container.StopOptions{})

	err = cli.ContainerKill(context.Background(), containerId, "")
	if err != nil {
		log.Fatal(err)
	}
}
