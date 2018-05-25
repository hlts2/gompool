package treiber

import (
	"errors"
	"sync"
)

var (

	// ErrStackEmpty represents error that stack is empty
	ErrStackEmpty = errors.New("stack is empty")

	// ErrStackFull represents error that stack is full
	ErrStackFull = errors.New("stack is full")
)

// Stack is treiber stack
type Stack struct {
	head *Node
	mu   *sync.Mutex
}

// NewStack returns stack instance
func NewStack() *Stack {
	return &Stack{
		head: nil,
		mu:   new(sync.Mutex),
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
	s.mu.Lock()

	if s.head == nil {
		newHead.next = nil
	} else {
		newHead.next = s.head
	}

	s.head = newHead

	s.mu.Unlock()
}

// Pop returns node of the stack
func (s *Stack) Pop() (*Node, error) {
	s.mu.Lock()

	if s.head == nil {
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next

	s.mu.Unlock()

	return tmpHead, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.head == nil
}
