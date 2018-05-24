package treiber

import (
	"testing"
)

func TestStackOperation(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	value, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", err)
	}

	if value != 3 {
		t.Errorf("Pop() value expected: %v, got: %v", 3, value)
	}

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() expected: %v, got: %v", true, stack.IsEmpty())
	}

	stack.Push(4)

	value, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error: %v", err)
	}

	if value != 4 {
		t.Errorf("Pop() value expected: %v, got: %v", 4, value)
	}

	value, err = stack.Pop()
	if err == nil {
		t.Error("Pop() error is nil")
	}

	if value != nil {
		t.Errorf("Pop() value expected: nil, got: %v", value)
	}
}
