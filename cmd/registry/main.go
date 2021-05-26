package main

import "github.com/busgo/elsa/internal/registry/server"

func main() {

	//var instance1 = &registry.Instance{
	//	Segment:         "dev",
	//	ServiceName:     "com.busgo.trade.proto.TradeService",
	//	Ip:              "192.168.1.1",
	//	Port:            8001,
	//	Metadata:        make(map[string]string),
	//	RegTimestamp:    time.Now().UnixNano(),
	//	UpTimestamp:     time.Now().UnixNano(),
	//	RenewTimestamp:  time.Now().UnixNano(),
	//	DirtyTimestamp:  time.Now().UnixNano(),
	//	LatestTimestamp: time.Now().UnixNano(),
	//}
	//r := registry.NewRegistry()
	//log.Printf("registry:%#v", r)
	//
	//in, err := r.Register(instance1)
	//if err != nil {
	//	log.Fatalf("register a instance fail:%#v", err)
	//}
	//log.Printf("register a instance success in:%#v", in)
	//
	//for {
	//
	//	time.Sleep(time.Second)
	//}


	rs := server.NewRegistryServer()
	if err := rs.Start();err !=nil{
		panic(err)
	}

}
