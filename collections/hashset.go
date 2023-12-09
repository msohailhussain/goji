package collections

import (
	"fmt"
)

type HashSet[T comparable] struct {
	m map[T]any
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{m: make(map[T]any)}
}

func (s *HashSet[T]) Add(element T) {
	_, exist := s.m[element]
	if !exist {
		s.m[element] = struct{}{}
	}
}

func (s *HashSet[T]) Remove(element T) {
	delete(s.m, element)
}

func (s HashSet[T]) Contains(element T) bool {
	_, exist := s.m[element]
	return exist
}

func (s HashSet[T]) ToSlice() []T {
	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s HashSet[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

func (s HashSet[T]) Len() int {
	return len(s.m)
}

func (s HashSet[T]) Union() int {
	return len(s.m)
}
