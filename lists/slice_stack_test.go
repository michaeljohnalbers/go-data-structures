package lists

import "testing"

func TestNewSliceStack(t *testing.T) {
	var initialCapacity uint64 = 10
	s := NewSliceStack(initialCapacity)
	if nil == s {
		t.Errorf("Failed to create new SliceStack: '%v'", s)
	}
	if len(s.stack) != 0 {
		t.Errorf("Incorrect initial size for SliceStack: %v", len(s.stack))
	}
	if uint64(cap(s.stack)) != initialCapacity {
		t.Errorf("Incorrect initial capacity for SliceStack: %v", cap(s.stack))
	}
	if s.initialCapacity != initialCapacity {
		t.Errorf("Incorrect initial capacity for SliceStack: %v", s.initialCapacity)
	}
}
