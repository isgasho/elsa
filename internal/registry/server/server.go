package server

import (
	"context"
	"github.com/busgo/elsa/internal/registry"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RegistryServer struct {
	s        *grpc.Server
	r        registry.Registry
	endpoint string
}

func NewRegistryServer(endpoint string) (*RegistryServer, error) {

	return &RegistryServer{
		s:        grpc.NewServer(),
		r:        registry.NewRegistry(),
		endpoint: endpoint,
	}, nil
}

// start registry server
func (rs *RegistryServer) Start() error {

	l, err := net.Listen("tcp", rs.endpoint)
	if err != nil {
		return err
	}
	pb.RegisterRegistryServiceServer(rs.s, rs)
	log.Printf("the registry server listen to:%s success...", rs.endpoint)
	if err = rs.s.Serve(l); err != nil {
		return err
	}

	return nil
}

// register a service instance
func (rs *RegistryServer) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

// fetch the instance with segment and service name
func (rs *RegistryServer) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchResponse, error) {
	return nil, nil
}

// renew the instance
func (rs *RegistryServer) Renew(ctx context.Context, in *pb.RenewRequest) (*pb.RenewResponse, error) {
	return nil, nil
}

// cancel the instance
func (rs *RegistryServer) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {
	return nil, nil
}
