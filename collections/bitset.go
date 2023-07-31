package collections

import (
	"math"
	"math/bits"
)

type Bitset[T any] struct {
	size             int
	startingBitIndex int // in [0,size)
	data             []uint
}

const arch = bits.UintSize

func NewBiset[T any](size int) *Bitset[T] {
	dim := (size-1)/arch + 1
	return &Bitset[T]{
		size:             size,
		startingBitIndex: 0,
		data:             make([]uint, dim),
	}
}
func (b *Bitset[T]) blockIndex(bitIndex int) int {
	return (bitIndex - 1) % arch
}

// zeroize all elements in [start, end)
// end can be before start
// 0 <= start < end < len(data)
func (b *Bitset[T]) zeroize(start, end int) {
	// middle part
	endBlock := b.blockIndex(end)
	block := (b.blockIndex(start) + 1) % len(b.data)
	start %= arch
	end %= arch
	if block == endBlock {
		// they are in same block
		b.data[block] &= (math.MaxUint << (arch - start)) & (math.MaxUint >> end)
		return
	}

	// left part
	b.data[block] &= math.MaxUint << (arch - start)

	// right part
	b.data[block] &= math.MaxUint >> end

	// middle part
	for ; block < endBlock; block = (block + 1) % len(b.data) {
		b.data[block] = 0
	}
}
func (b *Bitset[T]) zeroiseAll() {
	for i := range b.data {
		b.data[i] = 0
	}
}
func (b *Bitset[T]) ShiftLeft(n int) {
	if n >= b.size {
		b.zeroiseAll()
	}
	newStartingBitIndex := (b.startingBitIndex + n) % b.size
	b.zeroize(
		b.startingBitIndex,
		newStartingBitIndex,
	)
	b.startingBitIndex = newStartingBitIndex
}
func (b *Bitset[T]) ShiftRight(n int) {
	if n >= b.size {
		b.zeroiseAll()
	}
	wastedBits := (len(b.data) * arch) - b.size
	newStartingBitIndex := b.startingBitIndex - n
	startZeroize := newStartingBitIndex - wastedBits
	if startZeroize < 0 {
		startZeroize += b.size
		if newStartingBitIndex < 0 {
			newStartingBitIndex += b.size
		}
	}
	b.zeroize(
		startZeroize,
		b.startingBitIndex-wastedBits,
	)
	b.startingBitIndex = newStartingBitIndex
}

func (b *Bitset[T]) shitData(c int) {

}
