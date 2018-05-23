package treiber

import (
	"testing"
)

func TestStackOperation(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	ptr, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", ptr)
	}

	if ptr == nil {
		t.Error("Pop() ptr is nil")
	}

	if *ptr != 3 {
		t.Errorf("Pop() *ptr expected: %v, got: %v", 3, *ptr)
	}

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() expected: %v, got: %v", true, stack.IsEmpty())
	}

	stack.Push(4)

	ptr, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", err)
	}

	if ptr == nil {
		t.Error("Pop() ptr is nil")
	}

	if *ptr != 4 {
		t.Errorf("Pop() *ptr expected: %v, got: %v", 4, *ptr)
	}

	ptr, err = stack.Pop()
	if err == nil {
		t.Error("Pop() error is nil")
	}

	if ptr != nil {
		t.Error("Pop() ptr is not nil")
	}
}
