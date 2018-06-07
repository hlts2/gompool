package treiber

import (
	"errors"
	"sync/atomic"
)

var (

	// ErrStackEmpty represents error that stack is empty
	ErrStackEmpty = errors.New("stack is empty")
)

// Stack is treiber stack
type Stack struct {
	head *Node
	hasp int32
}

// NewStack returns stack instance
func NewStack() *Stack {
	return &Stack{
		head: nil,
		hasp: 0,
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
	for {
		if s.hasp == 0 && atomic.CompareAndSwapInt32(&s.hasp, 0, 1) {
			break
		}
	}

	newHead.next = s.head
	s.head = newHead

	s.hasp = 0
}

// Pop returns node of the stack
func (s *Stack) Pop() (*Node, error) {
	for {
		if s.hasp == 0 && atomic.CompareAndSwapInt32(&s.hasp, 0, 1) {
			break
		}
	}

	if s.head == nil {
		s.hasp = 0
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next

	s.hasp = 0

	return tmpHead, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *Stack) IsEmpty() bool {
	for {
		if s.hasp == 0 && atomic.CompareAndSwapInt32(&s.hasp, 0, 1) {
			break
		}
	}

	s.hasp = 0

	return s.head == nil
}

// Cap returns current capacity of stack
func (s *Stack) Cap() (cnt int) {
	for {
		if s.hasp == 0 && atomic.CompareAndSwapInt32(&s.hasp, 0, 1) {
			break
		}
	}

	tmpHead := s.head
	for tmpHead != nil {
		cnt++
		tmpHead = tmpHead.next
	}

	s.hasp = 0

	return cnt
}
