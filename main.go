package main

import (
	"bufio"
	"bytes"
	"fmt"
	"index/suffixarray"
	"log"
	"math"
	"net"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// When debug is assumed that input.txt file exists
func main() {
	io := IO{
		w: bufio.NewWriter(os.Stdout),
	}
	for _, arg := range os.Args {
		if arg == "∰" { // this should be passed only when debugging
			io.r = bufio.NewReader(First(os.Open("input.txt")))
			goto here
		}
	}
	io.r = bufio.NewReader(os.Stdin)
here:
	defer io.Flush()
	solve(&io)
}

// #region INTERFACES

// From https://pkg.go.dev/golang.org/x/exp/constraints
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface {
	Signed | Unsigned
}
type Float interface {
	~float32 | ~float64
}
type Complex interface {
	~complex64 | ~complex128
}
type Ordered interface {
	Integer | Float | ~string
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// #endregion
// #region TYPES WITH METHODS
// #region SingleLinkedList

// #endregion
// #region Queue

// #endregion
// #region Stack

// #endregion

// #region BinaryHeap

// #endregion

// #region Set

// #endregion

// #region TreeNode

// main source: https://github.com/shivamMg/ppds/blob/master/tree/tree.go

// #endregion

// #region Trie

// #endregion

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

// #endregion
// #endregion
// #region FUNCTIONS

// #endregion

// #region IO STUFF
type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func (io *IO) ScanInt8() (x int8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt16() (x int16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt32() (x int32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt() (x int)     { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanUInt8() (x uint8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt16() (x uint16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt32() (x uint32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt() (x uint)     { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanFloat32() (x float32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanFloat64() (x float64) { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanString() (x string) { fmt.Fscan(io.r, &x); return }

func (io *IO) Print(x ...any)   { fmt.Fprint(io.w, x...) }
func (io *IO) PrintLn(x ...any) { fmt.Fprintln(io.w, x...) }

func (io *IO) Flush() { io.w.Flush() }

// #endregion

// #region KEEP IMPORTS
func _() {
	_ = bufio.Reader{}
	_ = bytes.Buffer{}
	_ = suffixarray.Index{}
	_ = log.Default()
	_ = math.Abs(0)
	_ = net.Dialer{}
	_ = path.ErrBadPattern
	_ = sort.Float64sAreSorted(nil)
	_ = strconv.ErrRange
	_ = strings.Builder{}
	_ = time.ANSIC
}

// #endregion

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

func solve(io *IO) {
	for T := io.ScanUInt16(); T > 0; T-- {
		// SOLVE HERE
	}
}
