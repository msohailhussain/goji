package collections

import (
	"fmt"
)

type MultiSet[T comparable] struct {
	m map[T]int
}

func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{m: make(map[T]int)}
}

func (s *MultiSet[T]) Add(element T) {
	_, exist := s.m[element]
	if !exist {
		s.m[element] = 1
	}
}

func (s *MultiSet[T]) Remove(element T) {
	i, exist := s.m[element]
	if !exist || i == 1 {
		delete(s.m, element)
	} else {
		s.m[element] = i - 1
	}
}

func (s *MultiSet[T]) MultiplicityOf(element T) int {
	i, exist := s.m[element]
	if exist {
		return i
	} else {
		return 0
	}
}

func (s MultiSet[T]) Contains(element T) bool {
	_, exist := s.m[element]
	return exist
}

func (s MultiSet[T]) ToSlice() []T {
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

func (s MultiSet[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}
