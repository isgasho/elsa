package main

import "github.com/busgo/elsa/internal/registry/server"

func main() {
	r, _ := server.NewRegistryServer([]string{"127.0.0.1:8005"})

	if err := r.Start(); err != nil {
		panic(err)
	}
}
