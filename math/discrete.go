package math

import "golang.org/x/exp/constraints"

// At least one != 0
func GCD[T constraints.Unsigned](a, b T) T {
	if b < a {
		a, b = b, a
	}
	for {
		if a == 0 {
			return b
		}
		r := (b % a)
		b = a
		a = r
	}
}

// At least one != 0
func LCM[T constraints.Unsigned](a, b T) T {
	return (a * b) / GCD(a, b)
}
