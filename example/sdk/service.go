package main

import (
	"context"
	"github.com/busgo/elsa/pkg/client"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"runtime/pprof"
)

type PingGrpcService struct {
	serviceName string
}

func (s *PingGrpcService) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("ping request:%s", request.Ping)
	return &pb.PingResponse{
		Pong: "pong",
	}, nil
}

func main() {

	f, err := os.Create("cpu.prof")
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	stub, err := client.NewRegistryStub("dev", []string{"127.0.0.1:8005"})
	if err != nil {
		panic(err)
	}

	rpc := new(PingGrpcService)
	rpc.serviceName = "com.busgo.ping.proto.PingService"
	server := grpc.NewServer()
	pb.RegisterPingServiceServer(server, rpc)
	managed := client.NewManagedSentinel(8001, stub)

	managed.AddGrpcService(rpc.serviceName)

	l, err := net.Listen("tcp", ":8001")
	if err = server.Serve(l); err != nil {
		panic(err)
	}

}
