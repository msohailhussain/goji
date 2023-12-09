package collections

import (
	"fmt"
)

type MultiHashSet[T comparable] struct {
	m      map[T]int
	length int
}

func NewMultiHashSet[T comparable]() *MultiHashSet[T] {
	return &MultiHashSet[T]{m: make(map[T]int)}
}

func (s *MultiHashSet[T]) Add(element T) {
	_, exist := s.m[element]
	if !exist {
		s.m[element] = 1
	} else {
		s.m[element] = s.m[element] + 1
	}
	s.length++
}

func (s *MultiHashSet[T]) Remove(element T) {
	i, exist := s.m[element]
	if !exist || i == 1 {
		delete(s.m, element)
	} else {
		s.m[element] = i - 1
	}
	s.length--
}

func (s *MultiHashSet[T]) MultiplicityOf(element T) int {
	i, exist := s.m[element]
	if exist {
		return i
	} else {
		return 0
	}
}

func (s MultiHashSet[T]) Contains(element T) bool {
	_, exist := s.m[element]
	return exist
}

func (s MultiHashSet[T]) ToSlice() []T {
	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		i := s.m[k]
		for i > 0 {
			keys = append(keys, k)
			i--
		}
	}
	return keys
}

func (s MultiHashSet[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

func (s MultiHashSet[T]) Len() int {
	return s.length
}
