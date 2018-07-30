package treiber

import (
	"errors"

	lockfree "github.com/hlts2/lock-free"
)

var (

	// ErrStackEmpty represents error that stack is empty
	ErrStackEmpty = errors.New("stack is empty")
)

// Stack is treiber stack
type Stack struct {
	head *Node
	hasp int32
	lf   lockfree.LockFree
}

// NewStack returns stack instance
func NewStack() *Stack {
	return &Stack{
		head: nil,
		lf:   lockfree.New(),
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
	s.lf.Wait()

	newHead.next = s.head
	s.head = newHead

	s.lf.Signal()
}

// Pop returns node of the stack
func (s *Stack) Pop() (*Node, error) {
	s.lf.Wait()

	if s.head == nil {
		s.hasp = 0
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next

	s.lf.Signal()

	return tmpHead, nil
}

// IsEmpty returns true if the stack is empty, one the other hand, it returns false if it is not empty
func (s *Stack) IsEmpty() bool {
	defer s.lf.Signal()
	s.lf.Wait()
	return s.head == nil
}

// Cap returns current capacity of stack
func (s *Stack) Cap() (cnt int) {
	s.lf.Wait()

	tmpHead := s.head
	for tmpHead != nil {
		cnt++
		tmpHead = tmpHead.next
	}

	s.lf.Signal()

	return cnt
}
