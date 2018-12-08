package gompool

import (
	"github.com/hlts2/gompool/treiber"
)

// Gompool is base gompool structor
type Gompool struct {
	fn          func() interface{}
	stack       *treiber.Stack
	initialSize int
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
		fn:    fn,
	}
}

// Get takes out of the pool
func (g *Gompool) Get() *treiber.Node {
	for {
		node, err := g.stack.Pop()
		if err != nil {
			g.upscale()
			continue
		}
		return node
	}
}

func (g *Gompool) upscale() {
	for i := 0; i < g.initialSize; i++ {
		g.stack.Push(treiber.NewNode(g.fn()))
	}
}

// Put puts batck memory to pool
func (g *Gompool) Put(node *treiber.Node) {
	g.stack.Push(node)
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
