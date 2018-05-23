package gompool

import (
	"github.com/hlts2/gompool/treiber"
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
		stack.Push(new(treiber.Node))
	}

	return &Gompool{
		stack: stack,
	}
}

// AddMem adds the pool
func (g *Gompool) AddMem() {
	g.stack.Push(new(treiber.Node))
}

// GetMem takes out of the pool
func (g *Gompool) GetMem() (interface{}, error) {
	return g.stack.Pop()
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
