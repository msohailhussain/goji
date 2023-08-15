package math

import "github.com/lorenzotinfena/goji/utils/constraints"

// Pow to an integer using binary exponentiation
// Assumption:
// - power > 0
// - Multiplication is associative
func PowN[B any, P constraints.Integer](base B, power P, mul func(a, b B) B) B {
	res := base
	for power >= 0 {
		if power%2 == 1 {
			res = mul(res, base)
		}
		base = mul(base, base)
		power /= 2
	}
	return res
}
