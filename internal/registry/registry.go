package registry

import (
	"fmt"
	"log"
	"sync"
)

type Registry interface {

	// register a service instance
	Register(instance *Instance) (*Instance, error)

	// fetch instances with segment and service name
	Fetch(segment, serviceName string) ([]*Instance, error)

	// cancel a instance
	Cancel(segment, serviceName, ip string, port int32) (*Instance, error)

	// renew a instance
	Renew(segment, serviceName, ip string, port int32) (*Instance, error)
}

type registry struct {
	apps map[string]*Application
	sync.RWMutex
}

func NewRegistry() Registry {

	return &registry{
		apps:    make(map[string]*Application),
		RWMutex: sync.RWMutex{},
	}
}

// get app with segment and service name
func (r *registry) getApplication(segment, serviceName string) (*Application, bool) {
	r.RLock()
	defer r.RUnlock()
	app, ok := r.apps[fmt.Sprintf("%s-%s", segment, serviceName)]
	return app, ok
}

// register a service instance
func (r *registry) Register(instance *Instance) (*Instance, error) {
	log.Printf("start register action instance:%#v", instance)
	segment := instance.Segment
	serviceName := instance.ServiceName
	app, ok := r.getApplication(segment, serviceName)
	if !ok {
		app = NewApplication(segment, serviceName)
	}

	// add a new instance
	in, _ := app.addInstance(instance)
	if !ok {
		r.Lock()
		r.apps[fmt.Sprintf("%s-%s", segment, serviceName)] = app
		r.Unlock()
	}
	return in, nil
}

// fetch instances with segment service name
func (r *registry) Fetch(segment, serviceName string) ([]*Instance, error) {
	log.Printf("start fetch action segment:%s,serviceName:%s", segment, serviceName)
	app, ok := r.getApplication(segment, serviceName)
	if !ok {
		return nil, AppNotFoundError
	}
	ins := app.getAllInstance()
	return ins, nil
}

// cancel instance
func (r *registry) Cancel(segment, serviceName, ip string, port int32) (*Instance, error) {
	log.Printf("start cancel action segment:%s,serviceName:%s,ip:%s,port:%d", segment, serviceName, ip, port)
	app, ok := r.getApplication(segment, serviceName)
	if !ok {
		return nil, AppNotFoundError
	}
	in, ok := app.cancel(ip, port)
	return in, nil
}

// renew a instance
func (r *registry) Renew(segment, serviceName, ip string, port int32) (*Instance, error) {
	log.Printf("start renew action segment:%s,serviceName:%s,ip:%s,port:%d", segment, serviceName, ip, port)
	app, ok := r.getApplication(segment, serviceName)
	if !ok {
		return nil, AppNotFoundError
	}
	in, ok := app.renew(ip, port)
	if !ok {
		return nil, InstanceNotFoundError
	}
	return in, nil
}
