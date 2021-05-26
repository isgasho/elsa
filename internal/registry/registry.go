package registry

import (
	"fmt"
	"github.com/busgo/elsa/internal/registry/census"
	"log"
	"math/rand"
	"sync"
	"time"
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
	c *census.Census
}

func NewRegistry() Registry {

	r := &registry{
		apps:    make(map[string]*Application),
		c:       new(census.Census),
		RWMutex: sync.RWMutex{},
	}
	go r.evictLoop()
	return r
}

// get app with segment and service name
func (r *registry) getApplication(segment, serviceName string) (*Application, bool) {
	r.RLock()
	defer r.RUnlock()
	app, ok := r.apps[fmt.Sprintf("%s-%s", segment, serviceName)]
	return app, ok
}

// get all applications from apps
func (r *registry) getApplications() []*Application {
	r.RLock()
	defer r.RUnlock()
	if len(r.apps) == 0 {
		return make([]*Application, 0)
	}

	apps := make([]*Application, 0)

	for _, app := range r.apps {
		apps = append(apps, app)
	}
	return apps
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
	in, old := app.addInstance(instance)
	if !old {
		// increment need count
		r.c.IncrNeedCount()
	}
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
	if ok {
		// decrement the need renew count
		r.c.DecrNeedCount()
	}
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
	// increment renew count
	r.c.IncrCount()
	return in, nil
}

//----------------------------------------------scan evict task--------------------------------------------------//

// evict loop ticker
func (r *registry) evictLoop() {

	scanTicker := time.Tick(census.ScanEvictDuration)

	for {

		select {
		case <-scanTicker:
			r.c.ResetCount()
			r.evict()
		}
	}

}

// evict expired instance from apps
func (r *registry) evict() {
	log.Printf("evict task......")
	now := time.Now().UnixNano()
	apps := r.getApplications()
	if len(apps) == 0 {
		log.Printf("has no app to evict the instance...")
		return
	}

	var instanceSize int64
	var expiredSize int64
	protected := r.c.ProtectStatus()
	expiredInstances := make([]*Instance, 0)
	for _, app := range apps {

		ins := app.instances
		if len(ins) == 0 {
			continue
		}
		instanceSize += int64(len(ins))

		for _, in := range ins {

			// delta duration
			deltaDuration := now - in.RenewTimestamp

			// check expired instance
			if deltaDuration > int64(census.InstanceExpiredDuration) && !protected ||
				deltaDuration > int64(census.InstanceMaxExpiredDuration) {
				expiredInstances = append(expiredInstances, in)
			}

		}

		// check expired limit
		expiredSize = int64(len(expiredInstances))
		expiredLimit := instanceSize - int64(float64(instanceSize)*census.SelfProtectThreshold)
		if expiredSize > expiredLimit {
			expiredSize = expiredLimit
		}

		// no expired instance
		if expiredSize == 0 {
			log.Printf("has no expired instance....")
			return
		}

		// cancel the expired size instance
		for i := 0; i < int(expiredSize); i++ {
			j := i + rand.Intn(len(expiredInstances)-i)
			expiredInstances[i], expiredInstances[j] = expiredInstances[j], expiredInstances[i]

			expiredInstance := expiredInstances[i]
			in, err := r.Cancel(expiredInstance.Segment, expiredInstance.ServiceName, expiredInstance.Ip, expiredInstance.Port)
			if err != nil {
				log.Printf("cancel the expired instance fail:%#v instance:%#v", err, expiredInstance)
				continue
			}
			log.Printf("cancel the expired instance success instance:%#v", in)
		}

	}

}
