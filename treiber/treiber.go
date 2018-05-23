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
	value interface{}
}

// Push appends value into the stack
func (s *Stack) Push(value interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newHead := &Node{
		value: value,
	}

	if s.head == nil {
		newHead.next = nil
	} else {
		newHead.next = s.head
	}

	s.head = newHead

	return nil
}

// Pop returns item of the stack
func (s *Stack) Pop() (*interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.head == nil {
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next

	return &tmpHead.value, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.head == nil
}
