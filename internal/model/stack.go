package model

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	data []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, the second return value is false.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

// Len returns the number of items in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// IsEmpty returns true if the stack has no items.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Items returns items of stack
func (s *Stack[T]) Items() []T {
	return s.data
}
