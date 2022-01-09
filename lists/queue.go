package lists

// TODO: add copy?

// Queue represents a first-in, first-out (FIFO) data structure.
type Queue interface {
	// Clear removes all the elements from the queue
	Clear()
	// Enqueue adds the provided item to the queue. An error is returned if the item cannot be added.
	Enqueue(interface{}) error
	// Peek returns the first item in the queue, while leaving it in place. Null and an error is returned if the queue size is zero.
	Peek() (interface{}, error)
	// Pop removes and returns the first item in the queue. Null and an error is returned if the queue size is zero
	Pop() (interface{}, error)
	// Size returns the number of items in the queue
	Size() uint64
}
