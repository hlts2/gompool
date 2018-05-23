package gompool

import (
	"github.com/hlts2/gompool/treiber"
	"github.com/pkg/errors"
)

// Gompool is base gompool structor
type Gompool struct {
	stack *treiber.Stack
}

// NewGompool returns Gompool instance
func NewGompool(uSize uint) *Gompool {
	iSize := int(uSize)

	stack := treiber.NewStack()

	for i := 0; i < iSize; i++ {
		stack.Push(nil)
	}

	return &Gompool{
		stack: stack,
	}
}

// AddMem adds the pool
func (g *Gompool) AddMem() {
	g.stack.Push(nil)
}

// GetMem takes out of the pool
func (g *Gompool) GetMem() (*interface{}, error) {
	ptr, err := g.stack.Pop()
	if err != nil {
		return nil, errors.Wrap(err, "faild to get memory from pool")
	}

	return ptr, nil
}

// FreeMem puts batck memory to pool
func (g *Gompool) FreeMem(ptr *interface{}) {
	g.stack.Push(ptr)
}

// IsEmpty returns true if the pool is empty, one the other hand, it returns false if it is not empty
func (g *Gompool) IsEmpty() bool {
	return g.stack.IsEmpty()
}

// DestPool destroys all pools
func (g *Gompool) DestPool() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
