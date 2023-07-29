package sort

import (
	"github.com/lorenzotinfena/goji/utils/constraints"
)

func LowerBound[T constraints.Ordered](v []T, element T) int {
	l := 0
	r := len(v)
	for l != r {
		middle := (l + r) / 2
		if v[middle] < element {
			l = middle + 1
		} else {
			r = middle
		}
	}
	return l
}

func UpperBound[T constraints.Ordered](v []T, element T) int {
	l := 0
	r := len(v)
	for l != r {
		middle := (l + r) / 2
		if v[middle] <= element {
			l = middle + 1
		} else {
			r = middle
		}
	}
	return l
}
