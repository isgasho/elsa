package main

import (
	"context"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	c, err := grpc.Dial("127.0.0.1:8005", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewRegistryServiceClient(c)

	instance := &pb.ServiceInstance{
		Segment:         "dev",
		ServiceName:     "com.busgo.trade.proto.TradeService",
		Ip:              "192.168.1.1",
		Port:            8001,
		Metadata:        make(map[string]string),
		RegTimestamp:    time.Now().UnixNano(),
		UpTimestamp:     time.Now().UnixNano(),
		RenewTimestamp:  time.Now().UnixNano(),
		DirtyTimestamp:  time.Now().UnixNano(),
		LatestTimestamp: time.Now().UnixNano(),
	}
	resp, err := client.Register(context.Background(), &pb.RegisterRequest{
		Action:   pb.ActionType_Replication,
		Instance: instance,
	})

	if err != nil {
		panic(err)
	}
	log.Printf("response:%#v", resp)
}
