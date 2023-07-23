package math

import "golang.org/x/exp/constraints"

// Assumptions:
// - Both >= 0 and
// - At least one != 0
func GCD[T constraints.Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// Assumptions:
// - Both >= 0 and
// - At least one != 0
func LCM[T constraints.Unsigned](a, b T) T {
	return (a * b) / GCD(a, b)
}

// Assumptions:
// - Both >= 0 and
// - At least one != 0
// source: https://cp-algorithms.com/algebra/linear-diophantine-equation.html#algorithmic-solution
func ExtendedGCD[T constraints.Integer](a, b T) (T, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, xTmp, yTmp := ExtendedGCD(b, a%b)
	y := xTmp - yTmp*int(a/b)
	x := yTmp
	return gcd, x, y
}
