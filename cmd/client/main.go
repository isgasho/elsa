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

	instance2 := &pb.ServiceInstance{
		Segment:         "dev",
		ServiceName:     "com.busgo.trade.proto.TradeService",
		Ip:              "192.168.1.2",
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

	resp, err = client.Register(context.Background(), &pb.RegisterRequest{
		Action:   pb.ActionType_Replication,
		Instance: instance2,
	})

	if err != nil {
		panic(err)
	}
	log.Printf("response:%#v", resp)

	go func() {

		for {

			response, err := client.Fetch(context.Background(), &pb.FetchRequest{
				Segment:     "dev",
				ServiceName: "com.busgo.trade.proto.TradeService",
			})
			if err != nil {
				panic(err)
			}

			if response.Code == 0 {

				ins := response.Instances
				if len(ins) == 0 {
					log.Printf("the instance list is nil")
					break
				}
				for _, in := range ins {
					log.Printf("fetch service instance is :%#v", in)
				}
			}
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Minute * 60)

}
