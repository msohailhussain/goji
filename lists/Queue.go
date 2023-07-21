
package lists

type Queue[T comparable, I Unsigned] struct {
	l SingleLinkedList[T, I]
}

func NewQueue[T comparable, I Unsigned]() *Queue[T, I] {
	return &Queue[T, I]{
		l: SingleLinkedList[T, I]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (q *Queue[T, I]) Len() I          { return q.l.length }
func (q *Queue[T, I]) Enqueue(value T) { q.l.InsertLast(value) }
func (q *Queue[T, I]) Dequeue() T      { return q.l.RemoveFirst() }
func (q *Queue[T, I]) Preview() T      { return q.l.First() }
func (q Queue[T, I]) Log() string      { return q.l.Log() }
