package utils

import (
	"net"
)

const (
	Local = "127.0.0.1"
)

// get local ip
func LocalIPAddress() string {
	ins, err := net.Interfaces()
	if err != nil {
		return Local
	}
	for i := 0; i < len(ins); i++ {
		if (ins[i].Flags & net.FlagUp) != 0 {
			addresses, _ := ins[i].Addrs()

			for _, address := range addresses {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return Local
}
