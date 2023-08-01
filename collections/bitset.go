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
	block := (start%arch + 1) % len(b.data)
	endBlock := end % arch
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
func (b *Bitset[T]) ShiftLeft(c int) {
	if c >= b.size {
		b.zeroiseAll()
	}

	start := b.startingBitIndex
	b.startingBitIndex = b.startingBitIndex % b.size
	b.zeroize(start, b.startingBitIndex)
}
func (b *Bitset[T]) ShiftRight(c int) {
	if c >= b.size {
		b.zeroiseAll()
	}

	end := b.startingBitIndex
	b.startingBitIndex -= c
	if b.startingBitIndex < 0 {
		b.startingBitIndex += b.size
	}
	b.zeroize(b.startingBitIndex, end)
}

// Assumptions:
// - 0<c<arch
func (b *Bitset[T]) realCircularShiftLeft(c int) {
	var carry uint = 0
	for i := range b.data {
		carry, b.data[i] = b.data[i]<<(arch-c), carry|(b.data[i]>>c)
	}
	carry >>= (len(b.data) * arch) - b.size // wasted bits number
	carry |= b.data[len(b.data)-1] << uint(b.size) % arch
	b.data[0] |= carry
}

// Assumptions:
// - b.size == other.size
func (a *Bitset[T]) alignTo(b *Bitset[T]) {
	aIndex := a.startingBitIndex % a.size
	bIndex := b.startingBitIndex % a.size
	if aIndex != bIndex {
		
	}
}
