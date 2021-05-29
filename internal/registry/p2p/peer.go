package p2p

import (
	"context"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc"
	"log"
)

type SyncType int32

const (
	SyncRegType    SyncType = 1 // sync reg event
	SyncRenewType  SyncType = 2 // sync renew event
	syncCancelType SyncType = 3 // sync cancel event
)

// sync the registry message to other peer
type SyncMessage struct {
	Type    SyncType
	Content interface{}
}

// peer pool
type PeerPool struct {
	endpoints   []string
	peers       []*Peer
	syncMsgChan chan *SyncMessage
}

// peer with p2p
type Peer struct {
	endpoint string
	cli      pb.RegistryServiceClient
}

// new a peer pool
func NewPeerPool(endpoints []string) (*PeerPool, error) {

	peers := make([]*Peer, 0)
	for _, endpoint := range endpoints {
		p, err := NewPeer(endpoint)
		if err != nil {
			return nil, err
		}
		peers = append(peers, p)
	}
	pool := &PeerPool{
		endpoints:   endpoints,
		peers:       peers,
		syncMsgChan: make(chan *SyncMessage, 128),
	}
	go pool.lookup()
	return pool, nil
}

//  push a new sync message
func (pool *PeerPool) PushSyncMessage(message *SyncMessage) {
	pool.syncMsgChan <- message
}

// lookup the sync message chan
func (pool *PeerPool) lookup() {

	for {
		select {
		case message, ok := <-pool.syncMsgChan:
			if !ok {
				log.Printf("the sync message chan has closed...")
				return
			}
			pool.handleSyncMessage(message)
		}
	}
}

// handle sync message
func (pool *PeerPool) handleSyncMessage(message *SyncMessage) {

	if len(pool.peers) == 0 {
		return
	}
	switch message.Type {
	case SyncRegType:
		log.Printf("handle the reg sync message ....")
		pool.reg(message.Content.(*pb.RegisterRequest))
	case SyncRenewType:
		log.Printf("handle the renew sync message ....")
		pool.renew(message.Content.(*pb.RenewRequest))
	case syncCancelType:
		log.Printf("handle the cancel sync message ....")
		pool.cancel(message.Content.(*pb.CancelRequest))
	}

}

// sync register the service instance
func (pool *PeerPool) reg(request *pb.RegisterRequest) {

	for _, p := range pool.peers {
		in, err := p.cli.Register(context.Background(), request)
		if err != nil {
			log.Printf("the sync register service instance  fail:%#v", err)
			continue
		}
		log.Printf("the sync register service instance success :%#v", in)
	}
}

// sync renew the service instance
func (pool *PeerPool) renew(request *pb.RenewRequest) {

	for _, p := range pool.peers {
		in, err := p.cli.Renew(context.Background(), request)
		if err != nil {
			log.Printf("the sync renew service instance  fail:%#v", err)
			continue
		}
		log.Printf("the sync renew service instance success :%#v", in)
	}
}

// sync cancel the service instance
func (pool *PeerPool) cancel(request *pb.CancelRequest) {

	for _, p := range pool.peers {
		in, err := p.cli.Cancel(context.Background(), request)
		if err != nil {
			log.Printf("the sync cancel service instance  fail:%#v", err)
			continue
		}
		log.Printf("the sync cancel service instance success :%#v", in)
	}
}

// new a peer with endpoint
func NewPeer(endpoint string) (*Peer, error) {

	cc, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Peer{
		endpoint: endpoint,
		cli:      pb.NewRegistryServiceClient(cc),
	}, nil
}
