package set

import "testing"

func TestSetInt(t *testing.T) {
	set := NewSet[int]()
	set.Insert(1)
	if set.Len() != 1 {
		t.Error("set length should be 1")
	}
	set.Insert(2)
	set.Insert(3)

	if !set.Contains(2) {
		t.Error("set should contain 2")
	}

	if set.Contains(4) {
		t.Error("set should not contain 4")
	}

	set.Remove(2)

	if set.Contains(2) {
		t.Error("set should not contain 2")
	}

	if set.Len() != 2 {
		t.Error("set length should be 2")
	}

	if set.IsEmpty() {
		t.Error("set should not be empty")
	}
}
