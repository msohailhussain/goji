package collections

type Stack[T comparable] struct {
	l SingleLinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		l: SingleLinkedList[T]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (s *Stack[T]) Len() int      { return s.l.length }
func (s *Stack[T]) Push(value T)  { s.l.InsertFirst(value) }
func (s *Stack[T]) Pop() T        { return s.l.RemoveFirst() }
func (s *Stack[T]) Preview() T    { return s.l.First() }
func (s Stack[T]) String() string { return s.l.String() }
