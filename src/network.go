package src

import (
	"log"
	"net"
	"strings"
)

// get a list of all network interfaces
func GetNetworkInterfaces() []net.Interface {
	inter, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}

	return inter
}

// filter network interfaces so that you get only the one you are using for outside connections
func FilterNetworkInterfaces(interfaces *[]net.Interface) string {

	activeInterface := ""

	for _, iface := range *interfaces {
		if strings.Contains(iface.Flags.String(), "up") && !strings.Contains(iface.Flags.String(), "loopback") {
			addrs, err := iface.Addrs()

			if err != nil {
				log.Fatal(err)
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && ipNet.IP.To4()[0] != 172 {
					activeInterface = iface.Name
					break
				}
			}
		}
	}

	return activeInterface
}
