package main

import (
	"github.com/busgo/elsa/internal/registry"
	"log"
)

func main() {

	r := registry.NewRegistry()
	log.Printf("registry:%#v",r)
}
