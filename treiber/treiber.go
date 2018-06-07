package treiber

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

var (

	// ErrStackEmpty represents error that stack is empty
	ErrStackEmpty = errors.New("stack is empty")
)

// Stack is treiber stack
type Stack struct {
	head *Node
}

// NewStack returns stack instance
func NewStack() *Stack {
	return &Stack{
		head: nil,
	}
}

// Node is the item of stack
type Node struct {
	next  *Node
	Value interface{}
}

// NewNode returns Node instance
func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

// Push appends value into the stack
func (s *Stack) Push(newHead *Node) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&newHead.next)), unsafe.Pointer(s.head))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&s.head)), unsafe.Pointer(newHead))
}

// Pop returns node of the stack
func (s *Stack) Pop() (*Node, error) {
	tmpHead := (*Node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.head))))
	if tmpHead == nil {
		return nil, ErrStackEmpty
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&s.head)), unsafe.Pointer(tmpHead.next))

	return tmpHead, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *Stack) IsEmpty() bool {
	return atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.head))) == nil
}

// Cap returns current capacity of stack
func (s *Stack) Cap() (cnt int) {
	// TODO use aatomic

	tmpHead := s.head
	for tmpHead != nil {
		cnt++
		tmpHead = tmpHead.next
	}
	return cnt
}
