package client

import (
	"context"
	"github.com/busgo/elsa/pkg/utils"
	"log"
	"sync"
	"time"
)

type ManagedSentinel struct {
	sentinels map[string]*Sentinel
	stub      *RegistryStub
	sync.RWMutex
	ip   string
	port int32
}

// new a managed sentinel
func NewManagedSentinel(servicePort int32, registryStub *RegistryStub) *ManagedSentinel {
	return &ManagedSentinel{
		sentinels: make(map[string]*Sentinel),
		stub:      registryStub,
		RWMutex:   sync.RWMutex{},
		ip:        utils.LocalIPAddress(),
		port:      servicePort,
	}
}

// add grpc service
func (m *ManagedSentinel) AddGrpcService(serviceName string) {
	m.Lock()
	defer m.Unlock()

	sentinel, ok := m.sentinels[serviceName]
	if ok {
		return
	}
	sentinel = NewSentinel(serviceName, m.ip, m.port, m.stub)
	sentinel.register()
	m.sentinels[serviceName] = sentinel

}

// sentinel
type Sentinel struct {
	serviceName  string
	registryStub *RegistryStub
	ip           string
	port         int32
	closed       chan bool
	registerChan chan bool
}

// new a sentinel
func NewSentinel(serviceName, ip string, port int32, stub *RegistryStub) *Sentinel {

	s := &Sentinel{
		serviceName:  serviceName,
		registryStub: stub,
		ip:           ip,
		port:         port,
		closed:       make(chan bool),
		registerChan: make(chan bool, 1),
	}

	go s.lookup()
	return s
}

func (s *Sentinel) lookup() {

	ticker := time.Tick(time.Second * 30)
	for {
		select {
		case <-ticker:
			log.Printf("start renew the service name :%s", s.serviceName)
			s.renew()
		case <-s.registerChan: // register again the service instance
			time.Sleep(time.Second * 3)
			log.Printf("start register the service name :%s", s.serviceName)
			s.register()
		}
	}
}

func (s Sentinel) register() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	err := s.registryStub.Register(ctx, s.serviceName, s.ip, s.port)
	if err != nil {
		log.Printf("register to the service name :%s,ip:%s,port:%d,fail 3s after try again...", s.serviceName, s.ip, s.port)
		//s.registerChan <- true
	}
}

func (s *Sentinel) renew() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	state := s.registryStub.Renew(ctx, s.serviceName, s.ip, s.port)
	if !state {
		log.Printf("register to the service name :%s,ip:%s,port:%d,fail 3s after try again...", s.serviceName, s.ip, s.port)
		s.registerChan <- true
		return
	}
}

func (s *Sentinel) cancel() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	state := s.registryStub.Cancel(ctx, s.serviceName, s.ip, s.port)
	if !state {
		log.Printf("cancel to the service name :%s,ip:%s,port:%d", s.serviceName, s.ip, s.port)
	}
	s.closed <- true
}
