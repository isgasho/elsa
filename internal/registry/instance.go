package registry

import (
	"github.com/busgo/elsa/pkg/proto/pb"
	"time"
)

type Instance struct {
	Segment         string
	ServiceName     string
	Ip              string
	Port            int32
	Metadata        map[string]string
	RegTimestamp    int64
	UpTimestamp     int64
	RenewTimestamp  int64
	DirtyTimestamp  int64
	LatestTimestamp int64
}

// copy a new instance
func (instance *Instance) Copy() *Instance {
	in := new(Instance)
	*in = *instance
	return in
}

// new Instance
func NewInstance(instance *pb.ServiceInstance) *Instance {

	now := time.Now().UnixNano()
	return &Instance{
		Segment:         instance.Segment,
		ServiceName:     instance.ServiceName,
		Ip:              instance.Ip,
		Port:            instance.Port,
		Metadata:        instance.Metadata,
		RegTimestamp:    now,
		UpTimestamp:     now,
		RenewTimestamp:  now,
		DirtyTimestamp:  now,
		LatestTimestamp: now,
	}
}

// new a service instance
func NewServiceInstance(instance *Instance) *pb.ServiceInstance {

	return &pb.ServiceInstance{
		Segment:         instance.Segment,
		ServiceName:     instance.ServiceName,
		Ip:              instance.Ip,
		Port:            instance.Port,
		Metadata:        instance.Metadata,
		RegTimestamp:    instance.RegTimestamp,
		UpTimestamp:     instance.UpTimestamp,
		RenewTimestamp:  instance.RegTimestamp,
		DirtyTimestamp:  instance.DirtyTimestamp,
		LatestTimestamp: instance.LatestTimestamp,
	}
}
