package src

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

type FormattedTime struct {
	Hours   int32
	Minutes int32
	Seconds int32
}

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

// Cpu usage
func CPUData(perCPU bool) float64 {
	percent, err := cpu.Percent(time.Second, perCPU)
	if err != nil {
		log.Fatal(err)
	}

	return percent[0]
}

// Virtual Memory
func VirtualMemory() float64 {
	memory, err := mem.VirtualMemory()

	if err != nil {
		log.Fatal("Error while accessing virtual memory diagnostics")
	}

	return memory.UsedPercent
}

func FormatTime(seconds uint64) FormattedTime {
	time := FormattedTime{
		Hours:   0,
		Minutes: 0,
		Seconds: 0,
	}

	for seconds > 3600 {
		time.Hours++
		seconds -= 3600
	}

	for seconds > 60 {
		time.Minutes++
		seconds -= 60
	}

	time.Seconds = int32(seconds)
	return time
}

type HostInfoStats struct {
	Hostname        string
	Uptime          FormattedTime
	OS              string
	Platform        string
	PlatformFamily  string
	PlatformVersion string
	KernelVersion   string
	KernelArch      string
}

func HostInfo() interface{} {
	hostInfo, err := host.Info()

	if err != nil {
		log.Fatal("Error getting host data")
	}

	hostInfoFormatted := HostInfoStats{
		Hostname:        hostInfo.Hostname,
		Uptime:          FormatTime(hostInfo.Uptime),
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformFamily:  hostInfo.PlatformFamily,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		KernelArch:      hostInfo.KernelArch,
	}

	return hostInfoFormatted
}
