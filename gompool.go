package gompool

import (
	"github.com/hlts2/gompool/treiber"
)

// Gompool is base gompool structor
type Gompool struct {
	stack *treiber.Stack
}

// NewGompool returns Gompool instance
func NewGompool() *Gompool {
	return &Gompool{
		stack: treiber.NewStack(),
	}
}

// AddMem appends value into the pool
func (g *Gompool) AddMem(value interface{}) error {
	return g.stack.Push(value)
}

// GetMem takes out of value from the pool
func (g *Gompool) GetMem() (interface{}, error) {
	return g.stack.Pop()
}

// IsEmpty returns true if the pool is empty, one the other hand, it returns false if it is not empty
func (g *Gompool) IsEmpty() bool {
	return g.stack.IsEmpty()
}

// DestPool destroys all values of pool
func (g *Gompool) DestPool() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
