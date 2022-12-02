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
	"unicode/utf8"
)

func main() {
	io := newIO()
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

type Prioritizable[T any] interface {
	PriorTo(x T) bool // If you take this as a relation: if (a,b) in R => (b,a) not in R
}

// #endregion
// #region TYPES WITH METHODS
// #region SingleLinkedList
type singleLinkedListNode[T comparable] struct {
	Value T
	Next  *singleLinkedListNode[T]
}

type SingleLinkedList[T comparable, IndexType Unsigned] struct {
	first  *singleLinkedListNode[T]
	last   *singleLinkedListNode[T]
	length IndexType
}

func NewSingleLinkedList[T comparable, IndexType Unsigned]() *SingleLinkedList[T, IndexType] {
	return &SingleLinkedList[T, IndexType]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}
func (l *SingleLinkedList[T, I]) Len() I { return l.length }

func (l *SingleLinkedList[T, I]) First() T { return l.first.Value }

func (l *SingleLinkedList[T, I]) Last() T { return l.last.Value }

func (l *SingleLinkedList[T, I]) InsertFirst(value T) {
	if l.length == 0 {
		nodeToInsert := &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.first = &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
	}
	l.length++
}

func (l *SingleLinkedList[T, I]) InsertLast(value T) {
	if l.length == 0 {
		nodeToInsert := &singleLinkedListNode[T]{
			Value: value,
			Next:  l.first,
		}
		l.first = nodeToInsert
		l.last = nodeToInsert
	} else {
		l.last.Next = &singleLinkedListNode[T]{
			Value: value,
			Next:  nil,
		}
		l.last = l.last.Next
	}
	l.length++
}

// index <= length
func (l *SingleLinkedList[T, I]) InsertAt(index I, value T) {
	if index == 0 {
		l.InsertFirst(value)
		return
	}
	if index == l.length {
		l.InsertLast(value)
		return
	}

	n := l.first
	for index > 1 {
		n = n.Next
		index--
	}
	n.Next = &singleLinkedListNode[T]{
		Value: value,
		Next:  n.Next,
	}
	l.length++
}

func (l *SingleLinkedList[T, I]) Contains(value T) bool {
	tmp := l.first
	for i := I(0); i < l.length; i++ {
		if tmp.Value == value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}
func (l *SingleLinkedList[T, I]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// index < length
func (l *SingleLinkedList[T, I]) GetElementAt(index I) T {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	return n.Value
}

// index < length
func (l *SingleLinkedList[T, I]) SetElementAt(index I, value T) {
	n := l.first
	for index > 0 {
		n = n.Next
		index--
	}
	n.Value = value
}

func (l *SingleLinkedList[T, I]) RemoveFirst() T {
	tmp := l.first
	l.first = l.first.Next
	l.length--
	if l.length == 0 {
		l.last = nil
	}
	return tmp.Value
}
func (l *SingleLinkedList[T, I]) RemoveLast() (value T) {
	value = l.last.Value
	if l.length == 1 {
		l.first = nil
		l.last = nil
		l.length = 0
		return
	}

	tmp := l.first
	for i := I(2); i < l.length; i++ {
		tmp = tmp.Next
	}
	l.last = tmp
	return
}

// index < length
func (l *SingleLinkedList[T, I]) RemoveAt(index I) T {
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.length-1 {
		return l.RemoveLast()
	}

	tmp := l.first
	for i := I(1); i < index; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next.Value
	tmp.Next = tmp.Next.Next
	return res
}

func (l *SingleLinkedList[T, I]) ToSlice() (res []T) {
	res = make([]T, 0, l.length)
	tmp := l.first
	for i := I(0); i < l.length; i++ {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}
	return
}
func (it SingleLinkedList[T, I]) String() string {
	return "Qua:: " + fmt.Sprint(it.ToSlice())
}

// #endregion
// #region Queue

type Queue[T comparable, I Unsigned] struct {
	l SingleLinkedList[T, I]
}

func NewQueue[T comparable, I Unsigned]() *Queue[T, I] {
	return &Queue[T, I]{
		l: SingleLinkedList[T, I]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (q *Queue[T, I]) Len() I          { return q.l.length }
func (q *Queue[T, I]) Enqueue(value T) { q.l.InsertLast(value) }
func (q *Queue[T, I]) Dequeue() T      { return q.l.RemoveFirst() }
func (q *Queue[T, I]) Preview() T      { return q.l.First() }
func (q Queue[T, I]) String() string   { return q.l.String() }

// #endregion
// #region Stack

type Stack[T comparable, I Unsigned] struct {
	l SingleLinkedList[T, I]
}

func NewStack[T comparable, I Unsigned]() *Stack[T, I] {
	return &Stack[T, I]{
		l: SingleLinkedList[T, I]{
			first:  nil,
			last:   nil,
			length: 0,
		},
	}
}
func (s *Stack[T, I]) Len() I        { return s.l.length }
func (s *Stack[T, I]) Push(value T)  { s.l.InsertFirst(value) }
func (s *Stack[T, I]) Pop() T        { return s.l.RemoveFirst() }
func (s *Stack[T, I]) Preview() T    { return s.l.First() }
func (s Stack[T, I]) String() string { return s.l.String() }

// #endregion

// #region Heap
type BinaryHeap[T Prioritizable[T], I Signed] struct {
	s []T
}

func NewBinaryHeapFromSlice[T Prioritizable[T], I Signed](s []T) (h *BinaryHeap[T, I]) {
	h = &BinaryHeap[T, I]{
		s: s,
	}
	for i := (h.Len() - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return
}
func NewBinaryHeap[T Prioritizable[T], I Signed]() *BinaryHeap[T, I] {
	return &BinaryHeap[T, I]{s: make([]T, 0)}
}
func (h *BinaryHeap[T, I]) Len() I {
	return I(len(h.s))
}
func (h *BinaryHeap[T, I]) Push(value T) {
	h.s = append(h.s, value)
	h.heapifyUp(h.Len() - 1)
}
func (h *BinaryHeap[T, I]) Pop() (res T) {
	res = h.s[0]
	h.s[0] = h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeap[T, I]) heapifyDown(index I) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < h.Len() {
			if h.s[j-1].PriorTo(h.s[j]) {
				j--
			}
		} else {
			j--
			if j >= h.Len() {
				break
			}
		}
		if h.s[j].PriorTo(h.s[index]) {
			h.s[j], h.s[index] = h.s[index], h.s[j]
			index = j
		} else {
			break
		}
	}
	return origin != index
}
func (h *BinaryHeap[T, I]) heapifyUp(index I) {
	for {
		parent := (index - 1) / 2
		if parent == index || h.s[parent].PriorTo(h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		index = parent
	}
}
func (h *BinaryHeap[T, I]) Preview() T {
	return h.s[0]
}
func (h BinaryHeap[T, I]) String() string {
	return "" // #TODO
}

// #endregion

// #region Set
type Set[T comparable] struct {
	l SingleLinkedList[T, uint64]
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{l: *NewSingleLinkedList[T, uint64]()}
}
func (s *Set[T]) Add(element T) (added bool) {
	if !s.l.Contains(element) {
		s.l.InsertLast(element)
		return true
	} else {
		return false
	}
}
func (s *Set[T]) ToSlice() (res []T) {
	return s.l.ToSlice()
}
func (s *Set[T]) String() (res []T) {
	return s.l.ToSlice()
}

// #endregion
// #region TreeNode
type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

func (root TreeNode[T]) String() string {
	const (
		BoxVer       = "│"
		BoxHor       = "─"
		BoxVerRight  = "├"
		BoxDownLeft  = "┐"
		BoxDownRight = "┌"
		BoxDownHor   = "┬"
		BoxUpRight   = "└"
		Gutter       = 2
	)
	parents := map[*TreeNode[T]]*TreeNode[T]{}
	var setParents func(parents map[*TreeNode[T]]*TreeNode[T], root *TreeNode[T])
	setParents = func(parents map[*TreeNode[T]]*TreeNode[T], root *TreeNode[T]) {
		for _, c := range root.Children {
			parents[c] = root
			setParents(parents, c)
		}
	}
	setParents(parents, &root)

	var lines func(root *TreeNode[T]) (s []string)

	lines = func(root *TreeNode[T]) (s []string) {
		data := fmt.Sprintf("%s %v ", BoxHor, root.Value)
		l := len(root.Children)
		if l == 0 {
			s = append(s, data)
			return
		}

		w := utf8.RuneCountInString(data) // remove import unidcode/utf8
		for i, c := range root.Children {
			for j, line := range lines(c) {
				if i == 0 && j == 0 {
					if l == 1 {
						s = append(s, data+BoxHor+line)
					} else {
						s = append(s, data+BoxDownHor+line)
					}
					continue
				}

				var box string
				if i == l-1 && j == 0 {
					box = BoxUpRight
				} else if i == l-1 {
					box = " "
				} else if j == 0 {
					box = BoxVerRight
				} else {
					box = BoxVer
				}
				s = append(s, strings.Repeat(" ", w)+box+line)
			}
		}
		return
	}

	var lines2 func(root *TreeNode[T]) (s []string)
	lines2 = func(root *TreeNode[T]) (s []string) {
		s = append(s, fmt.Sprintf("%v", root.Value))
		l := len(root.Children)
		if l == 0 {
			return
		}

		for i, c := range root.Children {
			for j, line := range lines2(c) {
				// first line of the last child
				if i == l-1 && j == 0 {
					s = append(s, BoxUpRight+BoxHor+" "+line)
				} else if j == 0 {
					s = append(s, BoxVerRight+BoxHor+" "+line)
				} else if i == l-1 {
					s = append(s, "   "+line)
				} else {
					s = append(s, BoxVer+"  "+line)
				}
			}
		}
		return
	}
	safeData := func(n *TreeNode[T]) string {
		data := fmt.Sprintf("%v", n.Value)
		if data == "" {
			return " "
		}
		return data
	}
	var width func(widths map[*TreeNode[T]]int, n *TreeNode[T]) int
	width = func(widths map[*TreeNode[T]]int, n *TreeNode[T]) int {
		if w, ok := widths[n]; ok {
			return w
		}

		w := utf8.RuneCountInString(safeData(n)) + Gutter
		widths[n] = w
		if len(n.Children) == 0 {
			return w
		}

		sum := 0
		for _, c := range n.Children {
			sum += width(widths, c)
		}
		if sum > w {
			widths[n] = sum
			return sum
		}
		return w
	}

	var setPaddings func(paddings map[*TreeNode[T]]int, widths map[*TreeNode[T]]int, pad int, root *TreeNode[T])
	setPaddings = func(paddings map[*TreeNode[T]]int, widths map[*TreeNode[T]]int, pad int, root *TreeNode[T]) {
		for _, c := range root.Children {
			paddings[c] = pad
			setPaddings(paddings, widths, pad, c)
			pad += width(widths, c)
		}
	}

	isLeftMostChild := func(n *TreeNode[T]) bool {
		p, ok := parents[n]
		if !ok {
			// root
			return true
		}
		return p.Children[0] == n
	}

	paddings := map[*TreeNode[T]]int{}

	setPaddings(paddings, map[*TreeNode[T]]int{}, 0, &root)

	q := NewQueue[*TreeNode[T], uint64]()
	q.Enqueue(&root)
	linesss := []string{}
	for q.Len() > 0 {
		branches, nodes := "", ""
		covered := 0
		qLen := q.Len()
		for i := uint64(0); i < qLen; i++ {
			n := q.Dequeue()
			for _, c := range n.Children {
				q.Enqueue(c)
			}

			spaces := paddings[n] - covered
			data := safeData(n)
			nodes += strings.Repeat(" ", spaces) + data

			w := utf8.RuneCountInString(data)
			covered += spaces + w
			var preview *TreeNode[T] = nil
			if q.Len() > 0 {
				preview = q.Preview()
			}
			current, next := isLeftMostChild(n), isLeftMostChild(preview)
			if current {
				branches += strings.Repeat(" ", spaces)
			} else {
				branches += strings.Repeat(BoxHor, spaces)
			}

			if current && next {
				branches += BoxVer
			} else if current {
				branches += BoxVerRight
			} else if next {
				branches += BoxDownLeft
			} else {
				branches += BoxDownHor
			}

			if next {
				branches += strings.Repeat(" ", w-1)
			} else {
				branches += strings.Repeat(BoxHor, w-1)
			}
		}
		linesss = append(linesss, branches, nodes)
	}

	s := ""
	for _, line := range linesss[1:] {
		s += strings.TrimRight(line, " ") + "\n"

	}
	return s
}

// #endregion
// #region SqrtDecompositionSimple
/*
You can use SqrtDecompositionSimple only if:

Given a function $Query:E_1,...,E_n\rightarrow Q$

if $\exist unionQ:Q,Q\rightarrow Q$

s.t.

- $\forall n\in \N > 1, 1\le i<n, E_1,..., E_n\in E \\ query(E_1,..., E_n)=unionQ(query(E_1,..., E_i), query(E_{i+1},...,E_n))$

- (Only if you want use $update$ function)
$\forall n\in \N > 0, E_1,..., E_n\in E \\ query(E_1,...,E_{new},..., E_n)=updateQ(query(E_1,...,E_{old},...,E_n), indexof(E_{old}), E_{new})$
*/
type SqrtDecompositionSimple[E any, Q any] struct {
	querySingleElement func(element E) Q
	unionQ             func(q1 Q, q2 Q) Q
	updateQ            func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize uint64
}

// len(elements) > 0
func NewSqrtDecompositionSimple[E any, Q any](
	elements []E,
	querySingleElement func(element E) Q,
	unionQ func(q1 Q, q2 Q) Q,
	updateQ func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecompositionSimple[E, Q] {
	sqrtDec := &SqrtDecompositionSimple[E, Q]{
		querySingleElement: querySingleElement,
		unionQ:             unionQ,
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
			sqrtDec.blocks[i/blockSize] = sqrtDec.unionQ(sqrtDec.blocks[i/blockSize], sqrtDec.querySingleElement(elements[i]))
		}
	}
	sqrtDec.blockSize = blockSize
	return sqrtDec
}

// start < end (non included). Both are valid
func (s *SqrtDecompositionSimple[E, Q]) Query(start uint64, end uint64) Q {
	firstIndexNextBlock := ((start / s.blockSize) + 1) * s.blockSize
	q := s.querySingleElement(s.elements[start])
	if firstIndexNextBlock > end { // if in same block
		start++
		for start < end {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}
	} else {
		// left side
		start++
		for start < firstIndexNextBlock {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}

		//middle part
		endBlock := end / s.blockSize
		for i := firstIndexNextBlock / s.blockSize; i < endBlock; i++ {
			q = s.unionQ(q, s.blocks[i])
		}

		// right part
		for i := endBlock * s.blockSize; i < end; i++ {
			q = s.unionQ(q, s.querySingleElement(s.elements[i]))
		}
	}
	return q
}

// index is valid
func (s *SqrtDecompositionSimple[E, Q]) Update(index uint64, newElement E) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}

// #endregion
// #endregion
// #region FUNCTIONS
func Max[T Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}
func Min[T Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
func Abs[T Integer | Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// At least one != 0
func GCD[T Unsigned](a, b T) T {
	if b < a {
		Swap(&a, &b)
	}
	for {
		if a == 0 {
			return b
		}
		r := (b % a)
		b = a
		a = r
	}
}

// At least one != 0
func LCM[T Unsigned](a, b T) T {
	return (a * b) / GCD(a, b)
}

func First[T any](a T, _ any) T { return a }

// #endregion

// #region IO STUFF
type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func newIO() IO {
	return IO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

// Is assumed that input.txt file exists
func (io *IO) SetFileInput() {
	io.r = bufio.NewReader(First(os.Open("input.txt")))
}

func (io *IO) ScanInt8() (x int8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt16() (x int16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt32() (x int32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt64() (x int64) { fmt.Fscan(io.r, &x); return }

func (io *IO) ScanUInt8() (x uint8)   { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt16() (x uint16) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt32() (x uint32) { fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt64() (x uint64) { fmt.Fscan(io.r, &x); return }

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
	// io.SetFileInput() // Uncomment this while only when debugging

	for T := io.ScanUInt16(); T > 0; T-- {
		// SOLVE HERE
	}
}