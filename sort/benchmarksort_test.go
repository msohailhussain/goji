package sort_test

import (
	so "sort"
	"testing"

	"github.com/lorenzotinfena/goji/sort"
)

const n = 100000

func BenchmarkStandardSort(b *testing.B) {
	v := gen(n)
	so.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
}

func BenchmarkRadixSort(b *testing.B) {
	v := gen(n)
	sort.RadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
}

func BenchmarkRadixSortNaive(b *testing.B) {
	v := gen(n)
	sort.RadixSortNaive(v, func(a uint) []uint { return []uint{uint(a)} })
}

func BenchmarkInPlaceMSDRadixSort(b *testing.B) {
	v := gen(n)
	sort.InPlaceMSDRadixSort(v, func(a uint) []uint { return []uint{uint(a)} })
}
