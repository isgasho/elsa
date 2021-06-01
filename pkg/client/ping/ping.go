package ping

import (
	"github.com/busgo/elsa/pkg/client"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
)

type PingServiceStub struct {
	pb.PingServiceClient
}

// create a ping service stub
func NewPingServiceStub(resolver *client.ElsaNamingResolver) *PingServiceStub {

	cc, err := grpc.Dial("com.busgo.ping.proto.PingService", grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(resolver)))
	if err != nil {
		panic(err)
	}
	return &PingServiceStub{PingServiceClient: pb.NewPingServiceClient(cc)}
}
