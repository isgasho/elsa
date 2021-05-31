package registry

import (
	"context"
	"log"
	"testing"
	"time"
)

var endpoints = []string{"127.0.0.1:8005", "129.168.1.1:8005", "192.168.1.2:8005"}

func TestNewRegistryStub(t *testing.T) {

	stub, err := NewRegistryStub("dev", endpoints)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("registry stub is :%#v", *stub)
}

func TestRegistryStub_Register(t *testing.T) {

	stub, err := NewRegistryStub("dev", endpoints)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("registry stub is :%#v", *stub)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	err = stub.Register(ctx, "com.busgo.trade.proto.TradeService", "192.168.1.1", 8001)
	if err != nil {

		t.Fatalf("err:%#v", err)
	}

	log.Printf("the register action success")

}
