package src

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
	"log"
)

type FormattedTime struct {
	Hours   int32
	Minutes int32
	Seconds int32
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
