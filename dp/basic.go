package dp

import "fmt"

// Top-down or memoization approach to dp
//
// Usage example for fiboacci:
// dp := NewDP(
//
//	func(get func(int) int,
//		k int) int {
//		if k <= 1{
//			return k
//		}
//		return get(k-1)+get(k-2)
//	},
//
// )
// dp.Get(20)
type DP[Key comparable, Value any] struct {
	m          map[Key]Value
	recurrence func(get func(Key) Value, k Key) Value // Recurrence equation
}

func NewDP[Key comparable, Value any](recurrence func(get func(Key) Value, k Key) Value) *DP[Key, Value] {
	return &DP[Key, Value]{
		m:          make(map[Key]Value),
		recurrence: recurrence,
	}
}

func (dp *DP[Key, Value]) Get(k Key) Value {
	value, found := dp.m[k]
	if !found {
		value = dp.recurrence(dp.Get, k)
		dp.m[k] = value
	}
	return value
}
func (dp *DP[Key, Value]) Set(k Key, v Value) {
	dp.m[k] = v
}
func (dp *DP[Key, Value]) String() string {
	s := ""
	for k, v := range dp.m {
		s += fmt.Sprint(k) + " -> " + fmt.Sprintln(v)
	}
	return s
}
