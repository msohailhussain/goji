package collections

type Queue[T comparable] struct {
	l SingleLinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		l: SingleLinkedList[T]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (q *Queue[T]) Len() int        { return q.l.length }
func (q *Queue[T]) Enqueue(value T) { q.l.InsertLast(value) }
func (q *Queue[T]) Dequeue() T      { return q.l.RemoveFirst() }
func (q *Queue[T]) Preview() T      { return q.l.First() }
func (q Queue[T]) String() string   { return q.l.String() }
