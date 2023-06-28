package src

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

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
	err = cli.ContainerStop(context.Background(), containerId, container.StopOptions{})

	if err != nil {
		log.Fatal(err)
	}
}
