package client

import (
	"context"
	"fmt"
	"github.com/busgo/elsa/pkg/proto/pb"
	"google.golang.org/grpc/naming"
	"log"
	"strings"
	"sync"
	"time"
)

// direct name resolver   for registry
type DirectNameResolver struct {
	endpoints []string
	updates   []*naming.Update
	state     bool
}

func NewDirectNameResolver(endpoints []string) *DirectNameResolver {
	return &DirectNameResolver{endpoints: endpoints, updates: make([]*naming.Update, 0), state: false}
}

// Resolve creates a Watcher for target.
func (d *DirectNameResolver) Resolve(target string) (naming.Watcher, error) {
	return d, nil
}

// Next blocks until an update or error happens. It may return one or more
// updates. The first call should get the full set of the results. It should
// return an error if and only if Watcher cannot recover.
func (d *DirectNameResolver) Next() ([]*naming.Update, error) {

	if !d.state {
		d.state = true
		log.Printf("start init direct updates .....")
		for _, endpoint := range d.endpoints {
			log.Printf("the direct name resolver has add %s", endpoint)
			d.updates = append(d.updates, &naming.Update{
				Op:   naming.Add,
				Addr: strings.TrimSpace(endpoint),
			})
		}

	}

	return d.updates, nil
}

// Close closes the Watcher.
func (d *DirectNameResolver) Close() {
	log.Printf("the direct name resolver has close...")
	d.state = false
}

type ElsaNamingResolver struct {
	segment      string
	registryStub *RegistryStub
	watchers     map[string]*ElsaNamingWatcher
	sync.RWMutex
}

// new elsa naming resolver
func NewElsaNamingResolver(segment string, stub *RegistryStub) *ElsaNamingResolver {

	return &ElsaNamingResolver{
		segment:      segment,
		registryStub: stub,
		watchers:     make(map[string]*ElsaNamingWatcher),
		RWMutex:      sync.RWMutex{},
	}

}

// Resolve creates a Watcher for target.
func (resolver *ElsaNamingResolver) Resolve(target string) (naming.Watcher, error) {

	resolver.Lock()
	defer resolver.Unlock()
	watcher, ok := resolver.watchers[target]
	if ok {
		return watcher, nil
	}
	watcher = NewElsaNamingWatcher(resolver.segment, target, resolver.registryStub)
	resolver.watchers[target] = watcher
	return watcher, nil
}

type ElsaNamingWatcher struct {
	segment     string
	serviceName string
	stub        *RegistryStub
	updates     []*naming.Update
	sync.RWMutex
	closed chan bool
}

// new a elsa naming watcher
func NewElsaNamingWatcher(segment, serviceName string, stub *RegistryStub) *ElsaNamingWatcher {

	watcher := &ElsaNamingWatcher{
		segment:     segment,
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
	ticker := time.Tick(time.Second * 60)
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
