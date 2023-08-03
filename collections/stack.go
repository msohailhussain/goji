package collections

type Stack[T any] struct {
	l SingleLinkedList[T]
}

func NewStack[T any](equals func(T, T) bool) *Stack[T] {
	return &Stack[T]{
		l: *NewSingleLinkedList[T](equals),
	}
}
func (s *Stack[T]) Len() int      { return s.l.length }
func (s *Stack[T]) Push(value T)  { s.l.InsertFirst(value) }
func (s *Stack[T]) Pop() T        { return s.l.RemoveFirst() }
func (s *Stack[T]) Preview() T    { return s.l.First() }
func (s Stack[T]) String() string { return s.l.String() }
