package goji

import (
	"math"
)

// #region SqrtDecomposition
/*
You can use SqrtDecomposition only if:

Given a function $Query:E_1,...,E_n\rightarrow Q$

if $\exist unionQ:Q,Q\rightarrow Q$

s.t.

- $\forall n\in \N > 1, 1\le i<n, E_1,..., E_n\in E \\ query(E_1,..., E_n)=unionQ(query(E_1,..., E_i), query(E_{i+1},...,E_n))$

- (Only if you want use $update$ function)
$\forall n\in \N > 0, E_1,..., E_n\in E \\ query(E_1,...,E_{new},..., E_n)=updateQ(query(E_1,...,E_{old},...,E_n), indexof(E_{old}), E_{new})$
*/
type SqrtDecomposition[E any, Q any] struct {
	querySingleElement func(element E) Q
	mergeQ             func(q1 Q, q2 Q) Q
	updateQ            func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize uint64
}

// len(elements) > 0
func NewSqrtDecomposition[E any, Q any](
	elements []E,
	querySingleElement func(element E) Q,
	mergeQ func(q1 Q, q2 Q) Q,
	updateQ func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecomposition[E, Q] {
	sqrtDec := &SqrtDecomposition[E, Q]{
		querySingleElement: querySingleElement,
		mergeQ:             mergeQ,
		updateQ:            updateQ,
		elements:           elements,
	}
	sqrt := math.Sqrt(float64(len(sqrtDec.elements)))
	blockSize := uint64(sqrt)
	numBlocks := uint64(math.Ceil(float64(len(elements)) / float64(blockSize)))
	sqrtDec.blocks = make([]Q, numBlocks)
	for i := uint64(0); i < uint64(len(elements)); i++ {
		if i%blockSize == 0 {
			sqrtDec.blocks[i/blockSize] = sqrtDec.querySingleElement(elements[i])
		} else {
			sqrtDec.blocks[i/blockSize] = sqrtDec.mergeQ(sqrtDec.blocks[i/blockSize], sqrtDec.querySingleElement(elements[i]))
		}
	}
	sqrtDec.blockSize = blockSize
	return sqrtDec
}

// start < end (non included). Both are valid
func (s *SqrtDecomposition[E, Q]) Query(start uint64, end uint64) Q {
	firstIndexNextBlock := ((start / s.blockSize) + 1) * s.blockSize
	q := s.querySingleElement(s.elements[start])
	if firstIndexNextBlock > end { // if in same block
		start++
		for start < end {
			q = s.mergeQ(q, s.querySingleElement(s.elements[start]))
			start++
		}
	} else {
		// left side
		start++
		for start < firstIndexNextBlock {
			q = s.mergeQ(q, s.querySingleElement(s.elements[start]))
			start++
		}

		// middle part
		endBlock := end / s.blockSize
		for i := firstIndexNextBlock / s.blockSize; i < endBlock; i++ {
			q = s.mergeQ(q, s.blocks[i])
		}

		// right part
		for i := endBlock * s.blockSize; i < end; i++ {
			q = s.mergeQ(q, s.querySingleElement(s.elements[i]))
		}
	}
	return q
}

// index is valid
func (s *SqrtDecomposition[E, Q]) Update(index uint64, newElement E) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
