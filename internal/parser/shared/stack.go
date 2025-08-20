package shared

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	Data []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.Data = append(s.Data, v)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, the second return value is false.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.Data) == 0 {
		return zero, false
	}
	v := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return v, true
}

// Len returns the number of items in the stack.
func (s *Stack[T]) Len() int {
	return len(s.Data)
}

// IsEmpty returns true if the stack has no items.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Data) == 0
}
