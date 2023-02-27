package lib

// Generic stack

type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It is a run-time error to call Pop on an empty stack.
func (s *Stack[T]) Pop() T {
	var zero T
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = zero // to avoid memory leak
	s.data = s.data[:i]
	return res
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Data() []T {
	return s.data
}
