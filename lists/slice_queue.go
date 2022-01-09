package lists

import (
	"fmt"
)

// SliceQueue implements a Queue backed by a slice
type SliceQueue struct {
	queue           []interface{}
	initialCapacity uint64
}

// NewSliceQueue creates a new Queue object
func NewSliceQueue(initialCapacity uint64) *SliceQueue {
	q := SliceQueue{
		queue:           make([]interface{}, 0, initialCapacity),
		initialCapacity: initialCapacity,
	}
	return &q
}

var (
	_ Queue = (*SliceQueue)(nil)
)

func (q *SliceQueue) Clear() {
	q.queue = make([]interface{}, 0, q.initialCapacity)
}

func (q *SliceQueue) Enqueue(item interface{}) error {
	q.queue = append(q.queue, item)
	return nil
}

func (q *SliceQueue) Peek() (interface{}, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("cannot peek, queue size is zero")
	}
	return q.queue[0], nil
}

func (q *SliceQueue) Pop() (interface{}, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("cannot pop, queue size is zero")
	}
	first := q.queue[0]
	q.queue = q.queue[1:]
	return first, nil
}

func (q *SliceQueue) Size() uint64 {
	return uint64(len(q.queue))
}

func (q *SliceQueue) String() string {
	return fmt.Sprint(q.queue)
}
