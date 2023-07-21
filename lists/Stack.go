package lists

type Stack[T comparable, I Unsigned] struct {
	l SingleLinkedList[T, I]
}

func NewStack[T comparable, I Unsigned]() *Stack[T, I] {
	return &Stack[T, I]{
		l: SingleLinkedList[T, I]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (s *Stack[T, I]) Len() I       { return s.l.length }
func (s *Stack[T, I]) Push(value T) { s.l.InsertFirst(value) }
func (s *Stack[T, I]) Pop() T       { return s.l.RemoveFirst() }
func (s *Stack[T, I]) Preview() T   { return s.l.First() }
func (s Stack[T, I]) Log() string   { return s.l.Log() }
