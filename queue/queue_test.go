package queue

import "testing"

func TestQueueInt(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	if queue.Len() != 1 {
		t.Error("queue length should be 1")
	}
	queue.Push(2)
	queue.Push(3)

	if queue.Pop() != 1 {
		t.Error("queue front should be 1")
	}
	if queue.Front() != 2 {
		t.Error("queue front should be 2")
	}
	if queue.Back() != 3 {
		t.Error("queue back should be 3")
	}
	if queue.Len() != 2 {
		t.Error("queue length should be 2")
	}
}

func TestQueueString(t *testing.T) {
	queue := NewQueue[string]()
	queue.Push("hello")
	if queue.Len() != 1 {
		t.Error("queue length should be 1")
	}
	queue.Push("world")
	queue.Push("!")

	if queue.Pop() != "hello" {
		t.Error("queue front should be hello")
	}

	if queue.Pop() != "world" {
		t.Error("queue front should be world")
	}

}
