package lockfree

import (
	"sync/atomic"
	"time"
)

type LockFree interface {
	Wait()
	Signal()
}

// lockFree impl LockFree
type lockFree struct {
	flag int32
}

// New initializes a new instance of the LockFree
func New() LockFree {
	return &lockFree{
		flag: 0,
	}
}

// Wait waits until the shared counter is available. Update the sharing counter if available
func (g *lockFree) Wait() {
	for {
		if g.flag == 0 && atomic.CompareAndSwapInt32(&g.flag, 0, 1) {
			break
		}

		time.Sleep(1 * time.Nanosecond)
	}
}

// Signal signals termination of use of shared counter
func (g *lockFree) Signal() {
	g.flag = 0
}
