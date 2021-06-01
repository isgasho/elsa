package client

import (
	"testing"
)

func TestNewDirectNameResolver(t *testing.T) {

	nameResolver := NewDirectNameResolver([]string{"127.0.0.1:8005", "129.168.1.1:8005", "192.168.1.2:8005"})
	t.Logf("name resolver is :%#v", nameResolver)
}

func TestNewElsaNamingResolver(t *testing.T) {

	stub, err := NewRegistryStub("dev", []string{"127.0.0.1:8005", "129.168.1.1:8005", "192.168.1.2:8005"})
	if err != nil {
		t.Fatal(err)
	}

	resolver := NewElsaNamingResolver(stub)
	t.Logf("the elsa naming resolver is :%#v", resolver)
}
