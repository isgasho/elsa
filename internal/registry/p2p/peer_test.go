package p2p

import (
	"github.com/busgo/elsa/pkg/proto/pb"
	"testing"
	"time"
)

// test new peer pool
func TestNewPeerPool(t *testing.T) {

	pool, err := NewPeerPool([]string{"127.0.0.1:8005", "192.168.1.1:8005"})
	if err != nil {
		panic(err)
	}
	t.Logf("pool:%#v", pool)
	time.Sleep(time.Second)
}

// test push sync message
func TestPeerPool_PushSyncMessage(t *testing.T) {

	pool, err := NewPeerPool([]string{"127.0.0.1:8005", "192.168.1.1:8005"})
	if err != nil {
		panic(err)
	}
	t.Logf("pool:%#v", pool)
	time.Sleep(time.Second)

	pool.PushSyncMessage(&SyncMessage{
		Type: SyncRegType,
		Content: &pb.RegisterRequest{
			Action: pb.ActionType_Replication,
			Instance: &pb.ServiceInstance{
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
			},
		},
	})

	t.Logf("push the sync reg message success....")

	time.Sleep(time.Second)
}
