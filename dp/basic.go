package dp

func NewDP[K comparable, V any](computeValue func(dp *DP[K, V], key K) (value V)) *DP[K, V] {
	return &DP[K, V]{
		m:            make(map[K]V),
		computeValue: computeValue,
	}
}

type DP[K comparable, V any] struct {
	m            map[K]V
	computeValue func(dp *DP[K, V], key K) (value V)
}

func (dp *DP[K, V]) Get(key K) (value V) {
	value, found := dp.m[key]
	if !found {
		value = dp.computeValue(dp, key)
	}
	return
}
