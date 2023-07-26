package dp

// Top-down or memoization approach to dp
type DP[K comparable, V any] struct {
	m            map[K]V
	computeValue func(get func(key K) (value V), key K) (value V)
}

func NewDP[K comparable, V any](computeValue func(get func(key K) (value V), key K) (value V)) *DP[K, V] {
	return &DP[K, V]{
		m:            make(map[K]V),
		computeValue: computeValue,
	}
}

func (dp *DP[K, V]) Get(key K) (value V) {
	value, found := dp.m[key]
	if !found {
		value = dp.computeValue(dp.Get, key)
		dp.m[key] = value
	}
	return
}
