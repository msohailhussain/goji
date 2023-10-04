package collections

import (
	"fmt"
)

type Set[T comparable] struct {
	m map[T]any
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]any)}
}

func (s *Set[T]) Add(element T) {
	_, exist := s.m[element]
	if !exist {
		s.m[element] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	delete(s.m, element)
}

func (s Set[T]) Contains(element T) bool {
	_, exist := s.m[element]
	return exist
}

func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}
