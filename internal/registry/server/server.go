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
	server   *grpc.Server
	registry registry.Registry
	endpoint string
}

func NewRegistryServer(endpoint string) (*RegistryServer, error) {

	return &RegistryServer{
		server:   grpc.NewServer(),
		registry: registry.NewRegistry(),
		endpoint: endpoint,
	}, nil
}

// start registry server
func (r *RegistryServer) Start() error {

	l, err := net.Listen("tcp", r.endpoint)
	if err != nil {
		return err
	}
	pb.RegisterRegistryServiceServer(r.server, r)
	log.Printf("the registry server listen to:%s success...", r.endpoint)
	if err = r.server.Serve(l); err != nil {
		return err
	}

	return nil
}

// register a service instance
func (r *RegistryServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	instance := registry.NewInstance(request.Instance)

	in, _ := r.registry.Register(instance)
	return &pb.RegisterResponse{
		Code:     0,
		Message:  "success",
		Instance: registry.NewServiceInstance(in),
	}, nil
}

// fetch the instance with segment and service name
func (r *RegistryServer) Fetch(ctx context.Context, request *pb.FetchRequest) (*pb.FetchResponse, error) {

	instances, err := r.registry.Fetch(request.Segment, request.ServiceName)
	if err != nil {
		e := err.(*registry.RegistryError)
		log.Printf("fetch fail:%s", err.Error())
		return &pb.FetchResponse{
			Code:      e.Code,
			Message:   err.Error(),
			Instances: make([]*pb.ServiceInstance, 0),
		}, nil
	}

	ins := make([]*pb.ServiceInstance, 0)

	for _, in := range instances {
		ins = append(ins, registry.NewServiceInstance(in))
	}

	return &pb.FetchResponse{
		Code:      0,
		Message:   "success",
		Instances: ins,
	}, nil
}

// renew the instance
func (r *RegistryServer) Renew(ctx context.Context, request *pb.RenewRequest) (*pb.RenewResponse, error) {

	in, err := r.registry.Renew(request.Segment, request.ServiceName, request.Ip, request.Port)
	if err != nil {
		e := err.(*registry.RegistryError)
		return &pb.RenewResponse{
			Code:    e.Code,
			Message: e.Error(),
		}, err
	}
	return &pb.RenewResponse{
		Code:     0,
		Message:  "success",
		Instance: registry.NewServiceInstance(in),
	}, nil
}

// cancel the instance
func (r *RegistryServer) Cancel(ctx context.Context, request *pb.CancelRequest) (*pb.CancelResponse, error) {
	in, err := r.registry.Cancel(request.Segment, request.ServiceName, request.Ip, request.Port)
	if err != nil {
		e := err.(*registry.RegistryError)
		return &pb.CancelResponse{
			Code:    e.Code,
			Message: e.Error(),
		}, err
	}
	return &pb.CancelResponse{
		Code:     0,
		Message:  "success",
		Instance: registry.NewServiceInstance(in),
	}, nil
}
