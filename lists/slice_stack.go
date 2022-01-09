package lists

type SliceStack struct {
	stack           []interface{}
	initialCapacity uint64
}

func NewSliceStack(initialCapacity uint64) *SliceStack {
	s := SliceStack{
		stack:           make([]interface{}, 0, initialCapacity),
		initialCapacity: initialCapacity,
	}
	return &s
}

// TODO: finish implementing
