package client

import (
	"google.golang.org/grpc/naming"
	"log"
	"strings"
	"sync"
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
func NewElsaNamingResolver(registryServiceStub *RegistryStub) *ElsaNamingResolver {

	return &ElsaNamingResolver{
		segment:      registryServiceStub.GetSegment(),
		registryStub: registryServiceStub,
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
	watcher = NewElsaNamingWatcher(target, resolver.registryStub)
	resolver.watchers[target] = watcher
	return watcher, nil
}
