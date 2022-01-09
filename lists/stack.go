package lists

// Stack represents a last-in, first-out (LIFO) data structure.
type Stack interface {
	// Clear removes all the elements from the stack
	Clear()
	// Peek returns the top item in the stack, while leaving it in place. Null and an error is returned if the stack size is zero.
	Peek() (interface{}, error)
	// Pop removes and returns the top item in the stack. Null and an error is returned if the stack size is zero
	Pop() (interface{}, error)
	// Push adds the provided item to the stack. An error is returned if the item cannot be added.
	Push(interface{}) error
	// Size returns the number of items in the stack
	Size() uint64
}
