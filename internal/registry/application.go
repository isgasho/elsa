package registry

import (
	"fmt"
	"sync"
)

type Application struct {

	segment string
	serviceName string
	instances map[string]*Instance
	sync.RWMutex
}

// new a app with segment and service name
func NewApplication(segment,serviceName string)*Application  {

	return &Application{
		segment:     segment,
		serviceName: serviceName,
		instances:   make(map[string]*Instance),
		RWMutex:     sync.RWMutex{},
	}
}

// add a new instance
func (app *Application)addInstance(instance *Instance)(*Instance,bool)  {

	app.Lock()
	defer app.Unlock()
	ip:=instance.Ip
	port:=instance.Port
	in,ok:=app.instances[fmt.Sprintf("%s-%d",ip,port)]
	if ok{
		in.UpTimestamp = instance.UpTimestamp
		if in.DirtyTimestamp>instance.DirtyTimestamp {
			instance = in
		}
	}
	app.instances[fmt.Sprintf("%s-%d",ip,port)] = instance

	return instance.Copy(),ok
}

// get all instance
func (app *Application)getAllInstance()[]*Instance  {

	app.RLock()
	defer app.RUnlock()
	if len(app.instances) ==0 {
		return make([]*Instance,0)
	}

	ins:= make([]*Instance,0)

	for _,in:=range app.instances{
		ins = append(ins,in.Copy())
	}
	return ins
}

// cancel a instance
func (app *Application)cancel(ip string,port int32)(*Instance,bool)  {

	app.Lock()
	defer app.Unlock()
	in,ok:=app.instances[fmt.Sprintf("%s-%d",ip,port)]
	if !ok  {
		return nil,false
	}
	instance:= in.Copy()
	delete(app.instances,fmt.Sprintf("%s-%d",ip,port))
	return instance,true
}
