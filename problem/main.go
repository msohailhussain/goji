package main

/*
type Scalar interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32  | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~uintptr | ~float32 | ~float64
}
func max[T Scalar](a T, b T) T { 
	if a >= b { return a}
	return b
}*/
const MAXINT32 = 2147483647
const MININT32 = -2147483648
func Max(a, b int) int {
	if (a >= b) {return a}
	return b
}
func Min(a, b int) int {
	if (a <= b) {return a}
	return b
}

func main() {
	
}
