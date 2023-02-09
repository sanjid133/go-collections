package set

type Set[T comparable] interface {
	Insert(T)
	Remove(T)
	Len() int
	Contains(T) bool
	IsEmpty() bool
}

type set[T comparable] struct {
	keys map[T]struct{}
}

func NewSet[T comparable](initial ...T) Set[T] {
	s := &set[T]{keys: make(map[T]struct{})}
	for _, key := range initial {
		s.Insert(key)
	}
	return s
}

func (s *set[T]) Insert(key T) {
	s.keys[key] = struct{}{}
}

func (s *set[T]) Remove(key T) {
	delete(s.keys, key)
}

func (s *set[T]) Len() int {
	return len(s.keys)
}

func (s *set[T]) Contains(key T) bool {
	_, ok := s.keys[key]
	return ok
}

func (s *set[T]) IsEmpty() bool {
	return s.Len() == 0
}
