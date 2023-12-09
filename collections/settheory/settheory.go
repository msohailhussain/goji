package settheory

import "github.com/lorenzotinfena/goji/collections"

type Set[T comparable] interface {
	collections.HashSet[T] | collections.MultiHashSet[T]
	Add(T)
	Remove(T)
	Contains(T) bool
	ToSlice() []T
	Len() int
}

func Union[T comparable, S Set[T]](s1, s2 S) S {
	var res S
	switch (interface{}(s1)).(type) {
	case collections.HashSet[T]:
		res = (interface{}(*collections.NewHashSet[T]())).(S)

	case collections.MultiHashSet[T]:
		res = (interface{}(*collections.NewMultiHashSet[T]())).(S)
	}
	for _, el := range s1.ToSlice() {
		res.Add(el)
	}
	for _, el := range s2.ToSlice() {
		res.Add(el)
	}
	return res
}

func Intersection[T comparable, S Set[T]](s1, s2 S) S {
	if s2.Len() < s1.Len() {
		s1, s2 = s2, s1
	}
	var res S
	switch (interface{}(s1)).(type) {
	case collections.HashSet[T]:
		res = (interface{}(*collections.NewHashSet[T]())).(S)

	case collections.MultiHashSet[T]:
		res = (interface{}(*collections.NewMultiHashSet[T]())).(S)
	}
	for _, el := range s1.ToSlice() {
		if s2.Contains(el) {
			res.Add(el)
		}
	}
	return res
}
