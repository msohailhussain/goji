package collections

import (
	"fmt"

	"github.com/lorenzotinfena/goji/utils"
)

type singleLinkedListNode[T comparable] struct {
	Value T
	Next  *singleLinkedListNode[T]
}

type SingleLinkedList[T comparable] struct {
	first  *singleLinkedListNode[T]
	last   *singleLinkedListNode[T]
	length int
}

func NewSingleLinkedList[T comparable]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}
func (l *SingleLinkedList[T]) Len() int { return l.length }

func (l *SingleLinkedList[T]) First() T { return l.first.Value }

func (l *SingleLinkedList[T]) Last() T { return l.last.Value }

func (l *SingleLinkedList[T]) InsertFirst(value T) {
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

func (l *SingleLinkedList[T]) InsertLast(value T) {
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
func (l *SingleLinkedList[T]) InsertAt(index int, value T) {
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

func (l *SingleLinkedList[T]) Contains(value T) bool {
	tmp := l.first
	for i := 0; i < l.length; i++ {
		if tmp.Value == value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *SingleLinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// index < length
func (l *SingleLinkedList[T]) GetElementAt(index int) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *SingleLinkedList[T]) SetElementAt(index int, value T) {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	n.Value = value
}

func (l *SingleLinkedList[T]) RemoveFirst() T {
	tmp := l.first
	l.first = l.first.Next
	l.length--
	if l.first == nil {
		l.last = nil
	}
	return tmp.Value
}

func (l *SingleLinkedList[T]) RemoveLast() (value T) {
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
func (l *SingleLinkedList[T]) RemoveAt(index int) T {
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

func (l *SingleLinkedList[T]) ToSlice() (res []T) {
	res = make([]T, 0, l.length)
	tmp := l.first
	for i := 0; i < l.length; i++ {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}
	return
}
func (l *SingleLinkedList[T]) Iterate() utils.Iterator[T] {
	return newSingleLinkedListIterator(l.first)
}

func (it SingleLinkedList[T]) String() string {
	return fmt.Sprint(it.ToSlice())
}

type singleLinkedListIterator[T comparable] struct {
	current *singleLinkedListNode[T]
}

func newSingleLinkedListIterator[T comparable](current *singleLinkedListNode[T]) utils.Iterator[T] {
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
