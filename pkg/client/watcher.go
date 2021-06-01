package client

import (
	"context"
	"fmt"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc/naming"
	"log"
	"sync"
	"time"
)

const RefreshDuration = time.Second * 60 //  60s  clear cache and refresh the address list

type ElsaNamingWatcher struct {
	segment     string
	serviceName string
	stub        *RegistryStub
	updates     []*naming.Update
	sync.RWMutex
	closed chan bool
}

// new a elsa naming watcher
func NewElsaNamingWatcher(serviceName string, stub *RegistryStub) *ElsaNamingWatcher {

	watcher := &ElsaNamingWatcher{
		segment:     stub.GetSegment(),
		serviceName: serviceName,
		stub:        stub,
		updates:     make([]*naming.Update, 0),
		closed:      make(chan bool, 0),
		RWMutex:     sync.RWMutex{},
	}

	go watcher.lookup()

	return watcher
}

// Next blocks until an update or error happens. It may return one or more
// updates. The first call should get the full set of the results. It should
// return an error if and only if Watcher cannot recover.
func (w *ElsaNamingWatcher) Next() ([]*naming.Update, error) {
	return w.updates, nil
}

// Close closes the Watcher.
func (w *ElsaNamingWatcher) Close() {
	w.closed <- true
}

func (w *ElsaNamingWatcher) lookup() {

	// start refresh
	w.refresh()
	ticker := time.Tick(RefreshDuration)
	for {

		select {
		case <-ticker:
			w.refresh()
		case <-w.closed:
			log.Printf("the elsa naming watcher has closed service name:[%s]", w.serviceName)
			return
		}
	}

}

// refresh the service instance
func (w *ElsaNamingWatcher) refresh() {
	w.Lock()
	defer w.Unlock()
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	instances, err := w.stub.Fetch(ctx, w.serviceName)
	if err != nil {
		log.Printf("refresh service name :%s fail:%#v", w.serviceName, err)
		return
	}

	// refresh the service instance
	w.refreshInstance(instances)
}

// refresh the service instances
func (w *ElsaNamingWatcher) refreshInstance(instances []*pb.ServiceInstance) {
	if len(instances) == 0 {
		w.updates = make([]*naming.Update, 0)
		return
	}

	updates := make([]*naming.Update, 0)

	for _, instance := range instances {
		updates = append(updates, &naming.Update{
			Op:   naming.Add,
			Addr: fmt.Sprintf("%s:%d", instance.Ip, instance.Port),
		})
	}
	w.updates = updates
}
