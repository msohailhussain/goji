package utils

import (
	"github.com/lorenzotinfena/goji/utils/constraints"
)

func Prioritize[T constraints.Ordered]() func(a, b T) bool {
	return func(a, b T) bool {
		return a < b
	}
}

func Equalize[T constraints.Equalized]() func(a, b T) bool {
	return func(a, b T) bool {
		return a == b
	}
}
