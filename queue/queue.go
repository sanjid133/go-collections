package queue

type Queue[T comparable] interface {
	Push(T)
	Pop() T
	Len() int
	Front() T
	Back() T
	IsEmpty() bool
}

type queue[T comparable] struct {
	keys []T
}

func NewQueue[T comparable]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Push(t T) {
	q.keys = append(q.keys, t)
}

func (q *queue[T]) Pop() T {
	if q.Len() == 0 {
		panic("queue is empty")
	}
	t := q.keys[0]
	q.keys = q.keys[1:]
	return t
}

func (q *queue[T]) Len() int {
	return len(q.keys)
}

func (q *queue[T]) Front() T {
	if q.Len() == 0 {
		panic("queue is empty")
	}
	return q.keys[0]
}

func (q *queue[T]) Back() T {
	if q.Len() == 0 {
		panic("queue is empty")
	}
	return q.keys[q.Len()-1]
}

func (q *queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
