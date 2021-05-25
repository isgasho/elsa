package registry

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
