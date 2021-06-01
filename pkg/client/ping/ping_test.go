package ping

import (
	"context"
	"github.com/busgo/elsa/pkg/client"
	"github.com/busgo/elsa/pkg/proto/pb"
	"testing"
)

// test new ping service stub
func TestNewPingServiceStub(t *testing.T) {

	stub, err := client.NewRegistryStub("dev", []string{"127.0.0.1:8005"})
	if err != nil {
		panic(err)
	}
	resolver := client.NewElsaNamingResolver(stub)
	pingServiceStub := NewPingServiceStub(resolver)
	t.Logf("create ping service stub success %#v", pingServiceStub)
	response, err := pingServiceStub.Ping(context.Background(), &pb.PingRequest{
		Ping: "ping",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("call ping method success:%s", response.Pong)
}
