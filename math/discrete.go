package math

import "golang.org/x/exp/constraints"

// Euclidean algorithm
// It returns the positive solution
// Assumptions:
// - At least one != 0
func GCD[T constraints.Integer](a, b T) T {
	if b == 0 {
		return Abs(a)
	}
	return GCD(b, a%b)
}

// It returns the positive solution
// If one is 0, returns 0
// Assumptions:
// - Both != 0
func LCM[T constraints.Integer](a, b T) T {
	return Abs(a*b) / GCD(a, b)
}
