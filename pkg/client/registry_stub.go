package client

import (
	"context"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const RegistryServiceName = "com.busgo.registry.proto.Registry"

type RegistryStub struct {
	endpoints []string
	segment   string
	cli       pb.RegistryServiceClient
}

// new a registry stub
func NewRegistryStub(segment string, endpoints []string) (*RegistryStub, error) {

	cc, err := grpc.Dial(RegistryServiceName, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(NewDirectNameResolver(endpoints))))
	if err != nil {
		return nil, err
	}

	cli := pb.NewRegistryServiceClient(cc)
	return &RegistryStub{
		segment:   segment,
		endpoints: endpoints,
		cli:       cli,
	}, nil

}

// register a service instance
func (stub *RegistryStub) Register(ctx context.Context, serviceName, ip string, port int32) error {

	request := &pb.RegisterRequest{
		Type: pb.ReplicationType_Yes,
		Instance: &pb.ServiceInstance{
			Segment:         stub.segment,
			ServiceName:     serviceName,
			Ip:              ip,
			Port:            port,
			Metadata:        make(map[string]string),
			RegTimestamp:    time.Now().UnixNano(),
			UpTimestamp:     time.Now().UnixNano(),
			RenewTimestamp:  time.Now().UnixNano(),
			DirtyTimestamp:  time.Now().UnixNano(),
			LatestTimestamp: time.Now().UnixNano(),
		},
	}

	_, err := stub.cli.Register(ctx, request)
	if err != nil {
		log.Printf("register the service[segment:%s,serviceName:%s,ip:%s,port:%d] fail", stub.segment, serviceName, ip, port)
		return err
	}
	log.Printf("register the service[segment:%s,serviceName:%s,ip:%s,port:%d] success", stub.segment, serviceName, ip, port)
	return nil
}

// fetch the instance with segment and service name
func (stub *RegistryStub) Fetch(ctx context.Context, serviceName string) ([]*pb.ServiceInstance, error) {

	request := &pb.FetchRequest{
		Segment:     stub.segment,
		ServiceName: serviceName,
	}
	response, err := stub.cli.Fetch(ctx, request)
	if err != nil || response.Code != 0 {
		log.Printf("register the service[segment:%s,serviceName:%s] fail", stub.segment, serviceName)
		return make([]*pb.ServiceInstance, 0), nil
	}
	log.Printf("register the service[segment:%s,serviceName:%s,instances:%#v] success", stub.segment, serviceName, response.Instances)
	return response.Instances, nil
}

// renew the instance
func (stub *RegistryStub) Renew(ctx context.Context, serviceName, ip string, port int32) bool {
	request := &pb.RenewRequest{
		Segment:     stub.segment,
		ServiceName: serviceName,
		Ip:          ip,
		Port:        port,
	}
	response, err := stub.cli.Renew(ctx, request)
	if err != nil || response.Code != 0 {
		log.Printf("renew the service[segment:%s,serviceName:%s,ip:%s,port:%d] fail", stub.segment, serviceName, ip, port)
		return false
	}
	log.Printf("renew the service[segment:%s,serviceName:%s,ip:%s,port:%d] success", stub.segment, serviceName, ip, port)
	return true
}

// cancel the instance
func (stub *RegistryStub) Cancel(ctx context.Context, serviceName, ip string, port int32) bool {
	request := &pb.CancelRequest{
		Segment:     stub.segment,
		ServiceName: serviceName,
		Ip:          ip,
		Port:        port,
	}
	response, err := stub.cli.Cancel(ctx, request)
	if err != nil || response.Code != 0 {
		log.Printf("cancel the service[segment:%s,serviceName:%s,ip:%s,port:%d] fail", stub.segment, serviceName, ip, port)
		return false
	}
	log.Printf("cancel the service[segment:%s,serviceName:%s,ip:%s,port:%d] success", stub.segment, serviceName, ip, port)
	return true
}
