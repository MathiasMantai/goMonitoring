package src

import (
	"log"
	"strings"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func SanitizeContainer(containers []types.Container) []types.Container{


	length := 12
	for i := 0; i < len(containers); i++ {
		//shorten container id
		containers[i].ID = containers[i].ID[0:(length)]

		//replace backslash in container names
		for j:= 0; j < len(containers[i].Names); j++ {
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