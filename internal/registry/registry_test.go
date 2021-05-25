package registry

import (
	"testing"
	"time"
)

var instance1 = &Instance{
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
var instance2 = &Instance{
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

// register a service instance
func TestRegistry_Register(t *testing.T) {
	// new a registry
	r := NewRegistry()
	in, err := r.Register(instance1)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)
}

// fetch service instance
func TestRegistry_Fetch(t *testing.T) {
	// new a registry
	r := NewRegistry()
	in, err := r.Register(instance1)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)

	ins, err := r.Fetch("dev", "com.busgo.trade.proto.TradeService")
	if err != nil {
		t.Fatalf("fetch  instances fail:%#v", err)
	}

	for _, in := range ins {
		t.Logf("fetch the  instance :%#v", in)
	}

}

// cancel service instance
func TestRegistry_Cancel(t *testing.T) {
	// new a registry
	r := NewRegistry()
	in, err := r.Register(instance1)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)

	in, err = r.Register(instance2)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)

	ins, err := r.Fetch("dev", "com.busgo.trade.proto.TradeService")
	if err != nil {
		t.Fatalf("fetch  instances fail:%#v", err)
	}

	for _, in := range ins {
		t.Logf("fetch the  instance :%#v", in)
	}

	in, err = r.Cancel("dev", "com.busgo.trade.proto.TradeService", "192.168.1.1", 8001)
	if err != nil {
		t.Fatalf("cancel the instance fail:%#v", err)
	}

	t.Logf("cancel the instance success :%#v", in)

	ins, err = r.Fetch("dev", "com.busgo.trade.proto.TradeService")
	if err != nil {
		t.Fatalf("fetch  instances fail:%#v", err)
	}

	for _, in := range ins {
		t.Logf("fetch the  instance :%#v", in)
	}
}

// renew service instance
func TestRegistry_Renew(t *testing.T) {
	// new a registry
	r := NewRegistry()
	in, err := r.Register(instance1)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)

	in, err = r.Register(instance2)
	if err != nil {
		t.Fatalf("register a instance fail:%#v", err)
	}
	t.Logf("register a instance success in:%#v", in)

	ins, err := r.Fetch("dev", "com.busgo.trade.proto.TradeService")
	if err != nil {
		t.Fatalf("fetch  instances fail:%#v", err)
	}

	for _, in := range ins {
		t.Logf("fetch the  instance :%#v", in)
	}

	in, err = r.Renew("dev", "com.busgo.trade.proto.TradeService", "192.168.1.1", 8001)
	if err != nil {
		t.Fatalf("renew instance fail :%#v", err)
	}
	t.Logf("renew instance success %#v", in)
}
