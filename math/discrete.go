
// At least one != 0
func GCD[T Unsigned](a, b T) T {
	if b < a {
		Swap(&a, &b)
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
func LCM[T Unsigned](a, b T) T {
	return (a * b) / GCD(a, b)
}