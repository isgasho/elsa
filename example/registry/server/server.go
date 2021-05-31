package main

import "github.com/busgo/elsa/internal/registry/server"

func main() {
	rs, _ := server.NewRegistryServer([]string{"127.0.0.1:8005", "192.168.1.1:8005"})

	if err := rs.Start(); err != nil {
		panic(err)
	}
}
