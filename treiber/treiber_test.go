package treiber

import (
	"testing"
)

func TestStackOperation(t *testing.T) {
	stack := NewStack()

	stack.Push(NewNode(1))
	stack.Push(NewNode(2))
	stack.Push(NewNode(3))

	node, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", err)
	}

	if node == nil {
		t.Error("Pop() node is nil")
	}

	if node.Value != 3 {
		t.Errorf("Pop() node.value expected: %v, got: %v", 3, node.Value)
	}

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() expected: %v, got: %v", true, stack.IsEmpty())
	}

	stack.Push(NewNode(4))

	node, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", err)
	}

	if node == nil {
		t.Error("Pop() node is nil")
	}

	if node.Value != 4 {
		t.Errorf("Pop() node.Value expected: %v, got: %v", 4, node.Value)
	}

	node, err = stack.Pop()
	if err == nil {
		t.Error("Pop() error is nil")
	}

	if node != nil {
		t.Errorf("Pop() node expected: nil, got: %v", node)
	}
}
