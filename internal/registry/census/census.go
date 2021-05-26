package census

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	RenewDuration        = 30 * time.Second
	ScanEvictDuration    = 60 * time.Second
	SelfProtectThreshold = 0.85
)

type Census struct {
	count      int64
	needCount  int64
	leastCount int64
	threshold  int64
	sync.RWMutex
}

func (c *Census) IncrCount() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Census) ResetCount() {
	atomic.StoreInt64(&c.leastCount, atomic.SwapInt64(&c.count, 0))
}

func (c *Census) IncrNeedCount() {
	c.Lock()
	defer c.Unlock()
	c.needCount += int64(float64(ScanEvictDuration) / float64(RenewDuration))
	c.threshold = int64(float64(c.needCount) * SelfProtectThreshold)
}

func (c *Census) DecrNeedCount() {
	c.Lock()
	defer c.Unlock()
	c.needCount -= int64(float64(ScanEvictDuration) / float64(RenewDuration))
	c.threshold = int64(float64(c.needCount) * SelfProtectThreshold)
}

func (c *Census) ProtectStatus() bool {
	return atomic.LoadInt64(&c.leastCount) < atomic.LoadInt64(&c.threshold)
}
