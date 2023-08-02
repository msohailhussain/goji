package collections

import (
	"math"
	"math/bits"
)

type Bitset struct {
	size int
	data []uint
}

const blockSize = bits.UintSize

func NewBiset(size int) *Bitset {
	dim := (size-1)/blockSize + 1
	return &Bitset{
		size: size,
		data: make([]uint, dim),
	}
}

func (b *Bitset) ShiftLeft(c int) {
	if c >= b.size {
		for i := range b.data {
			b.data[i] = 0
		}
		return
	}
	blockShift := c / blockSize
	for i := 0; i <= len(b.data)-1-blockShift; i++ {
		b.data[i] = b.data[i+blockShift]
	}
	for i := len(b.data) - blockShift; i < len(b.data); i++ {
		b.data[i] = 0
	}

	c %= blockSize

	var carry uint = 0
	for i := len(b.data)-1; i>= 0; i-- {
		carry, b.data[i] = b.data[i]>>(blockSize-c), carry|(b.data[i]<<c)
	}

}
func (b *Bitset) ShiftRight(c int) {
	if c >= b.size {
		for i := range b.data {
			b.data[i] = 0
		}
		return
	}
	blockShift := c / blockSize
	for i := len(b.data) - 1 - blockShift; i >= 0; i-- {
		b.data[i+blockShift] = b.data[i]
	}
	for i := 0; i < blockShift; i++ {
		b.data[i] = 0
	}

	c %= blockSize

	var carry uint = 0
	for i := range b.data {
		carry, b.data[i] = b.data[i]<<(blockSize-c), carry|(b.data[i]>>c)
	}
	b.data[len(b.data)-1] &= math.MaxUint << (len(b.data)*blockSize - b.size)
}

func (b *Bitset) Set(index int, value bool) {
	elem := &b.data[index/blockSize]
	index = index % blockSize
	tmp := uint(1) << (blockSize - index - 1)
	if value {
		*elem = *elem | tmp
	} else {
		*elem = *elem & ^tmp
	}
}
func (b *Bitset) Get(index int) bool {
	return b.data[index/blockSize]&(uint(1)<<(blockSize-index%blockSize-1)) != 0
}

func (b *Bitset) ToSlice() []bool {
	s := make([]bool, b.size)
	for i := 0; i < b.size; i++ {
		s[i] = b.Get(i)
	}
	return s
}

func (b *Bitset) Len() int {
	return b.size
}
