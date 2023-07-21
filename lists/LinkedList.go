package lists

type singleLinkedListNode[T comparable] struct {
	Value T
	Next  *singleLinkedListNode[T]
}
type singleLinkedListIterator[T comparable] struct {
	current *singleLinkedListNode[T]
}

func newSingleLinkedListIterator[T comparable](current *singleLinkedListNode[T]) Iterator[T] {
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

type SingleLinkedList[T comparable, IndexType Unsigned] struct {
	first  *singleLinkedListNode[T]
	last   *singleLinkedListNode[T]
	length IndexType
}

func NewSingleLinkedList[T comparable, IndexType Unsigned]() *SingleLinkedList[T, IndexType] {
	return &SingleLinkedList[T, IndexType]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}
func (l *SingleLinkedList[T, I]) Len() I { return l.length }

func (l *SingleLinkedList[T, I]) First() T { return l.first.Value }

func (l *SingleLinkedList[T, I]) Last() T { return l.last.Value }

func (l *SingleLinkedList[T, I]) InsertFirst(value T) {
	if l.length == 0 {
		nodeToInsert := &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.first = &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
	}
	l.length++
}

func (l *SingleLinkedList[T, I]) InsertLast(value T) {
	if l.length == 0 {
		nodeToInsert := &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.last.Next = &singleLinkedListNode[T]{
			Value: value,
			Next:  nil,
		}
		l.last = l.last.Next
	}
	l.length++
}

// index <= length
func (l *SingleLinkedList[T, I]) InsertAt(index I, value T) {
	if index == 0 {
		l.InsertFirst(value)
		return
	}
	if index == l.length {
		l.InsertLast(value)
		return
	}

	n := l.first
	for index > 1 {
		n = n.Next
		index--
	}
	n.Next = &singleLinkedListNode[T]{
		Value: value,
		Next:  n.Next,
	}
	l.length++
}

func (l *SingleLinkedList[T, I]) Contains(value T) bool {
	tmp := l.first
	for i := I(0); i < l.length; i++ {
		if tmp.Value == value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *SingleLinkedList[T, I]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// index < length
func (l *SingleLinkedList[T, I]) GetElementAt(index I) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *SingleLinkedList[T, I]) SetElementAt(index I, value T) {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	n.Value = value
}

func (l *SingleLinkedList[T, I]) RemoveFirst() T {
	tmp := l.first
	l.first = l.first.Next
	l.length--
	if l.first == nil {
		l.last = nil
	}
	return tmp.Value
}

func (l *SingleLinkedList[T, I]) RemoveLast() (value T) {
	value = l.last.Value
	if l.length == 1 {
		l.first = nil
		l.last = nil
		l.length = 0
		return
	}

	tmp := l.first
	for i := I(2); i < l.length; i++ {
		tmp = tmp.Next
	}
	l.last = tmp
	l.length--
	return
}

// index < length
func (l *SingleLinkedList[T, I]) RemoveAt(index I) T {
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.length-1 {
		return l.RemoveLast()
	}

	tmp := l.first
	for i := I(1); i < index; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next.Value
	tmp.Next = tmp.Next.Next
	l.length--
	return res
}

func (l *SingleLinkedList[T, I]) ToSlice() (res []T) {
	res = make([]T, 0, l.length)
	tmp := l.first
	for i := I(0); i < l.length; i++ {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}
	return
}

func (l *SingleLinkedList[T, I]) GetIterator() Iterator[T] {
	return newSingleLinkedListIterator(l.first)
}

func (it SingleLinkedList[T, I]) Log() string {
	return fmt.Sprint(it.ToSlice())
}