package shared

type Stack[T any] struct {
	Data []T
}

func (s *Stack[T]) Push(v T) {
	s.Data = append(s.Data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Data) == 0 {
		var zero T
		return zero, false
	}
	top := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return top, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Data) == 0
}
