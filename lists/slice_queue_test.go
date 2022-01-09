package lists

import (
	"testing"
)

func TestNewSliceQueue(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)
	if nil == q {
		t.Errorf("Failed to create new SliceQueue: '%v'", q)
	}
	if len(q.queue) != 0 {
		t.Errorf("Incorrect initial size for SliceQueue: %v", len(q.queue))
	}
	if uint64(cap(q.queue)) != initialCapacity {
		t.Errorf("Incorrect initial capacity for SliceQueue: %v", cap(q.queue))
	}
	if q.initialCapacity != initialCapacity {
		t.Errorf("Incorrect initial capacity for SliceQueue: %v", q.initialCapacity)
	}
}

func TestSliceQueue_Clear(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)

	if len(q.queue) != 2 {
		t.Errorf("Setup failed, invalid size: %v", len(q.queue))
	}

	q.Clear()

	if len(q.queue) != 0 {
		t.Errorf("Incorrect size after clear: %v", len(q.queue))
	}

	if uint64(cap(q.queue)) != initialCapacity {
		t.Errorf("Incorrect capacity after clear: %v", cap(q.queue))
	}
}

func TestSliceQueue_Enqueue(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	for ii := uint64(1); ii <= initialCapacity*2; ii++ {
		err := q.Enqueue(ii)
		if err != nil {
			t.Errorf("Enqueue failed on item %v. Error: '%v'", ii, err)
		}

		if uint64(len(q.queue)) != ii {
			t.Errorf("Incorrect queue size on item %v. Size: %v", ii, len(q.queue))
		}

		if q.queue[ii-1] != ii {
			t.Errorf("Invalid storage of item %v, index %v is %v", ii, ii-1, q.queue[ii-1])
		}
	}
}

func TestSliceQueue_Peek(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	item, err := q.Peek()

	if err == nil {
		t.Errorf("Unexpectedly returned nill error for empty peek.")
	}

	if item != nil {
		t.Errorf("Unexpectedly retunred a non-nil item for empty peek: '%v'", item)
	}

	_ = q.Enqueue("hello")
	_ = q.Enqueue("there")

	item, err = q.Peek()

	if err != nil {
		t.Errorf("Unexpected error returned from peek: %v", err)
	}

	if item != "hello" {
		t.Errorf("Incorrect item returned from peek: '%v'", item)
	}
}

func TestSliceQueue_Pop(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	item, err := q.Pop()
	if err == nil {
		t.Errorf("Unexpectedly returned nill error for empty pop.")
	}

	if item != nil {
		t.Errorf("Unexpectedly retunred a non-nil item for empty pop: '%v'", item)
	}

	_ = q.Enqueue("hello")
	_ = q.Enqueue("there")

	item, err = q.Pop()

	if err != nil {
		t.Errorf("Unexpected error returned from pop: %v", err)
	}

	if item != "hello" {
		t.Errorf("Incorrect item returned from pop: '%v'", item)
	}

	if len(q.queue) != 1 {
		t.Errorf("Incorrect queue size after pop: %v", len(q.queue))
	}

	if q.queue[0] != "there" {
		t.Errorf("Incorrect new front item: '%v'", q.queue[0])
	}
}

func TestSliceQueue_Size(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	if q.Size() != 0 {
		t.Errorf("Incorrect initial size: %v", q.Size())
	}

	_ = q.Enqueue(1)
	_ = q.Enqueue(1)
	_ = q.Enqueue(1)

	if q.Size() != 3 {
		t.Errorf("Incorrect size after adding items: %v", q.Size())
	}

	q.Clear()

	if q.Size() != 0 {
		t.Errorf("Incorrect size after clear: %v", q.Size())
	}
}

func TestSliceQueue_String(t *testing.T) {
	var initialCapacity uint64 = 10
	q := NewSliceQueue(initialCapacity)

	_ = q.Enqueue("hi there")
	_ = q.Enqueue(1)
	_ = q.Enqueue(3.3)

	s := q.String()

	if s != "[hi there 1 3.3]" {
		t.Errorf("Incorrect string conversion, '%v'", s)
	}
}
