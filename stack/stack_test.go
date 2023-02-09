package stack

import "testing"

func TestStackInt(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	if stack.Len() != 1 {
		t.Error("stack length should be 1")
	}
	stack.Push(2)
	stack.Push(3)

	if stack.Pop() != 3 {
		t.Error("stack top should be 3")
	}

	if stack.Top() != 2 {
		t.Error("stack top should be 2")
	}

	if stack.Len() != 2 {
		t.Error("stack length should be 2")
	}
}

func TestStackString(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("hello")
	if stack.Len() != 1 {
		t.Error("stack length should be 1")
	}
	stack.Push("world")
	stack.Push("!")

	if stack.Pop() != "!" {
		t.Error("stack top should be !")
	}
}
