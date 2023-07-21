func Max[T Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Abs[T Integer | Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}