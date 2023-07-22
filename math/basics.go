package math

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Diff[T constraints.Integer | constraints.Float](a, b T) T {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
