package collections

type circularDoublyLinkedListNode[T any] struct {
	Value      T
	Prev, Next *circularDoublyLinkedListNode[T]
}

type CircularDoublyLinkedList[T any] struct {
	first  *circularDoublyLinkedListNode[T]
	length int
	equals func(T, T) bool
}

// equals can be nil
func NewCircularDoublyLinkedList[T any](equals func(T, T) bool) *CircularDoublyLinkedList[T] {
	return &CircularDoublyLinkedList[T]{
		first:  nil,
		length: 0,
		equals: equals,
	}
}
func (l *CircularDoublyLinkedList[T]) Len() int { return l.length }

func (l *CircularDoublyLinkedList[T]) First() T { return l.first.Value }

func (l *CircularDoublyLinkedList[T]) Last() T { return l.first.Prev.Value }

func (l *CircularDoublyLinkedList[T]) InsertFirst(value T) {
	l.InsertLast(value)
	l.first = l.first.Prev
}

func (l *CircularDoublyLinkedList[T]) InsertLast(value T) {
	if l.length == 0 {
		l.first = &circularDoublyLinkedListNode[T]{
			Value: value,
			Prev:  l.first.Prev,
			Next:  l.first,
		}
		return
	}

	node := &circularDoublyLinkedListNode[T]{
		Value: value,
		Prev:  l.first.Prev,
		Next:  l.first,
	}
	node.Prev.Next = node
	l.first.Prev = node
	l.length++
}

// merge another ll after the last element
func (l *CircularDoublyLinkedList[T]) MergeEnd(ll *CircularDoublyLinkedList[T]) {
	if l.Len() == 0 {
		*l = *ll
		return
	}
	if ll.Len() == 0 {
		return
	}
	l.first.Prev.Next = ll.first
	ll.first.Prev.Next = l.first
	l.first.Prev, ll.first.Prev = ll.first.Prev, l.first.Prev
}

/* TODO
// index <= length
func (l *CircularDoublyLinkedList[T]) InsertAt(index int, value T) {
}

// Assumptions:
// - equals != nil
func (l *CircularDoublyLinkedList[T]) Contains(value T) bool {
	tmp := l.first
	for i := 0; i < l.length; i++ {
		if l.equals(tmp.Value, value) {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *CircularDoublyLinkedList[T]) GetElementEqualsTo(value T) (T, bool) {
	tmp := l.first
	for i := 0; i < l.length; i++ {
		if l.equals(tmp.Value, value) {
			return tmp.Value, true
		}
		tmp = tmp.Next
	}
	return value, false
}

func (l *CircularDoublyLinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// index < length
func (l *CircularDoublyLinkedList[T]) GetElementAt(index int) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *CircularDoublyLinkedList[T]) SetElementAt(index int, value T) {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	n.Value = value
}

func (l *CircularDoublyLinkedList[T]) RemoveFirst() T {
	tmp := l.first
	l.first = l.first.Next
	l.length--
	if l.first == nil {
		l.last = nil
	}
	return tmp.Value
}

func (l *CircularDoublyLinkedList[T]) RemoveLast() (value T) {
	value = l.last.Value
	if l.length == 1 {
		l.first = nil
		l.last = nil
		l.length = 0
		return
	}

	tmp := l.first
	for i := 2; i < l.length; i++ {
		tmp = tmp.Next
	}
	l.last = tmp
	l.length--
	return
}

// index < length
func (l *CircularDoublyLinkedList[T]) RemoveAt(index int) T {
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.length-1 {
		return l.RemoveLast()
	}

	tmp := l.first
	for i := 1; i < index; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next.Value
	tmp.Next = tmp.Next.Next
	l.length--
	return res
}

func (l *CircularDoublyLinkedList[T]) Remove(element T) bool {
	if l.Len() == 0 {
		return false
	}

	if l.equals(l.First(), element) {
		l.RemoveFirst()
		return true
	}
	if l.equals(l.Last(), element) {
		l.RemoveLast()
		return true
	}
	tmp := l.first
	for i := 2; i < l.Len(); i++ {
		if l.equals(tmp.Next.Value, element) {
			tmp.Next = tmp.Next.Next
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *CircularDoublyLinkedList[T]) ToSlice() (res []T) {
	res = make([]T, 0, l.length)
	tmp := l.first
	for i := 0; i < l.length; i++ {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}
	return
}

func (l *SinglyLinkedList[T]) GetIterator() utils.Iterator[T] {
	return newSingleLinkedListIterator(l.first)
}

func (it SinglyLinkedList[T]) String() string {
	return fmt.Sprint(it.ToSlice())
}

type singleLinkedListIterator[T any] struct {
	current *singlyLinkedListNode[T]
}

func newSingleLinkedListIterator[T any](current *singlyLinkedListNode[T]) utils.Iterator[T] {
	return &singleLinkedListIterator[T]{
		current: current,
	}
}

func (it *singleLinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *singleLinkedListIterator[T]) Next() T {
	tmp := it.current.Value
	it.current = it.current.Next
	return tmp
}
*/
