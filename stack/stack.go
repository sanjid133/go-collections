package stack

type Stack[T comparable] interface {
	Push(T)
	Pop() T
	Top() T
	Len() int
	IsEmpty() bool
}

type stack[T comparable] struct {
	keys []T
}

func NewStack[T comparable]() Stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(key T) {
	s.keys = append(s.keys, key)
}

func (s *stack[T]) Pop() T {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	t := s.keys[s.Len()-1]
	s.keys = s.keys[:s.Len()-1]
	return t
}

func (s *stack[T]) Top() T {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	return s.keys[s.Len()-1]
}

func (s *stack[T]) Len() int {
	return len(s.keys)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
