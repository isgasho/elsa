package main

import "github.com/busgo/elsa/internal/registry/server"

func main() {
	rs, _ := server.NewRegistryServer("127.0.0.1:8005")

	if err := rs.Start(); err != nil {
		panic(err)
	}

}
