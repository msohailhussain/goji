package dp

// Top-down or memoization approach to dp
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

func (dp *DP[Key, Value]) Get(key Key) Value {
	value, found := dp.m[key]
	if !found {
		value = dp.recurrence(dp.Get, key)
		dp.m[key] = value
	}
	return value
}
