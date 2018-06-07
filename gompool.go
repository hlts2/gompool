package gompool

import (
	"github.com/hlts2/gompool/treiber"
	"github.com/pkg/errors"
)

// Gompool is base gompool structor
type Gompool struct {
	fn    func() interface{}
	stack *treiber.Stack
}

// NewGompool returns Gompool instance
func NewGompool(uSize uint, fn func() interface{}) *Gompool {
	iSize := int(uSize)

	stack := treiber.NewStack()

	for i := 0; i < iSize; i++ {
		stack.Push(treiber.NewNode(fn()))
	}

	return &Gompool{
		stack: stack,
	}
}

// Add adds the pool
func (g *Gompool) Add() {
	g.stack.Push(treiber.NewNode(g.fn()))
}

// Get takes out of the pool
func (g *Gompool) Get() (*treiber.Node, error) {
	node, err := g.stack.Pop()
	if err != nil {
		return nil, errors.Wrap(err, "faild to get memory from pool")
	}

	return node, nil
}

// Put puts batck memory to pool
func (g *Gompool) Put(node *treiber.Node) {
	g.stack.Push(node)
}

// IsEmpty returns true if the pool is empty, one the other hand, it returns false if it is not empty
func (g *Gompool) IsEmpty() bool {
	return g.stack.IsEmpty()
}

// Cap returns current capacity of pool
func (g *Gompool) Cap() int {
	return g.stack.Cap()
}

// DestPool destroys all pools
func (g *Gompool) DestPool() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
