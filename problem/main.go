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

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// #endregion
// #region TYPES WITH METHODS
// #region SingleLinkedList
type singleLinkedListNode[T comparable] struct {
	Value T
	Next  *singleLinkedListNode[T]
}
type singleLinkedListIterator[T comparable] struct {
	current *singleLinkedListNode[T]
}

func newSingleLinkedListIterator[T comparable](current *singleLinkedListNode[T]) Iterator[T] {
	return &singleLinkedListIterator[T]{
		current: current,
	}
}
func (it *singleLinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}
func (it *singleLinkedListIterator[T]) Next() T {
	tmp := it.current.Value
	it.current = it.current.Next
	return tmp
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
	if l.first == nil {
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
	l.length--
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
	l.length--
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

func (l *SingleLinkedList[T, I]) GetIterator() Iterator[T] {
	return newSingleLinkedListIterator(l.first)
}
func (it SingleLinkedList[T, I]) String() string {
	return fmt.Sprint(it.ToSlice())
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

// #region BinaryHeap
type BinaryHeap[T any] struct {
	s     []T
	prior func(T, T) bool
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapFromSlice[T any](
	s []T,
	Prior func(T, T) bool) (h *BinaryHeap[T]) {
	h = &BinaryHeap[T]{
		s:     s,
		prior: Prior,
	}
	if h.Len() > 1 {
		for i := (int64(h.Len()) - 2) / 2; i >= 0; i-- {
			h.heapifyDown(int64(i))
		}
	}
	return
}
func NewBinaryHeap[T any](Prior func(T, T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{s: make([]T, 0), prior: Prior}
}
func (h *BinaryHeap[T]) Len() uint64 {
	return uint64(len(h.s))
}
func (h *BinaryHeap[T]) Push(value T) {
	h.s = append(h.s, value)
	h.heapifyUp(int64(h.Len() - 1))
}
func (h *BinaryHeap[T]) Pop() (res T) {
	res = h.s[0]
	h.s[0] = h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeap[T]) heapifyDown(index int64) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < int64(h.Len()) {
			if h.prior(h.s[j-1], h.s[j]) {
				j--
			}
		} else {
			j--
			if j >= int64(h.Len()) {
				break
			}
		}
		if h.prior(h.s[j], h.s[index]) {
			h.s[j], h.s[index] = h.s[index], h.s[j]
			index = j
		} else {
			break
		}
	}
	return origin != index
}
func (h *BinaryHeap[T]) heapifyUp(index int64) {
	for {
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.s[parent], h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		index = parent
	}
}
func (h *BinaryHeap[T]) Preview() T {
	return h.s[0]
}
func (h BinaryHeap[T]) String() string {
	return "" // #TODO
}

// #endregion

// #region Set
type Set[T comparable] struct {
	m map[T]any
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}
func (s *Set[T]) Add(element T) {
	_, exist := s.m[element]
	if !exist {
		s.m[element] = struct{}{}
	}
}
func (s *Set[T]) Contains(element T) bool {
	_, exist := s.m[element]
	return exist
}
func (s *Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}
func (s *Set[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

// #endregion

// #region TreeNode
type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

// main source: https://github.com/shivamMg/ppds/blob/master/tree/tree.go
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
	mergeQ             func(q1 Q, q2 Q) Q
	updateQ            func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize uint64
}

// len(elements) > 0
func NewSqrtDecompositionSimple[E any, Q any](
	elements []E,
	querySingleElement func(element E) Q,
	mergeQ func(q1 Q, q2 Q) Q,
	updateQ func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecompositionSimple[E, Q] {
	sqrtDec := &SqrtDecompositionSimple[E, Q]{
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
func (s *SqrtDecompositionSimple[E, Q]) Query(start uint64, end uint64) Q {
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

		//middle part
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

// Use this if in SqrtDecompsitionSimple, the mergeQ function is hard, but an expandQ function is easy
func MoSAlgorithm[E any, Q any](
	querySingleElement func(element E) Q,
	expandQ func(*Q, E),
	clone func(Q) Q,
	elements []E,
	queries []struct {
		left  uint64
		right uint64
	},
) (res []Q) {
	// Initialization
	res = make([]Q, len(queries))
	sqrt := math.Sqrt(float64(len(elements)))
	blockSize := uint64(sqrt)
	blocks := make([]*SingleLinkedList[struct {
		left  uint64
		right uint64
		index uint64
	}, uint64], uint64(math.Ceil(float64(len(elements))/float64(blockSize))))
	for i := range blocks {
		blocks[i] = NewSingleLinkedList[struct {
			left  uint64
			right uint64
			index uint64
		}, uint64]()
	}
	for i, v := range queries {
		blocks[v.left/blockSize].InsertLast(struct {
			left  uint64
			right uint64
			index uint64
		}{
			left:  v.left,
			right: v.right,
			index: uint64(i),
		})
	}
	blockSorted := make([]*Queue[struct {
		left  uint64
		right uint64
		index uint64
	}, uint64], len(blocks))
	for i := range blocks {
		tmp := blocks[i].ToSlice()
		SelectionSort(tmp, func(a, b struct {
			left  uint64
			right uint64
			index uint64
		}) bool {
			return a.right < b.right
		})
		blockSorted[i] = NewQueue[struct {
			left  uint64
			right uint64
			index uint64
		}, uint64]()
		for _, tmp2 := range tmp {
			blockSorted[i].Enqueue(tmp2)
		}
	}
	// Main
	for i := uint64(0); i < uint64(len(blockSorted)); i++ {
		block := blockSorted[i]
		if block.Len() == 0 {
			continue
		}
		for block.Len() > 0 {
			if block.Preview().right/blockSize == i {
				q := block.Dequeue()
				res[q.index] = querySingleElement(elements[q.left])
				q.left++
				for q.left < q.right {
					expandQ(&res[q.index], elements[q.left])
				}
			} else {
				break
			}
		}
		middleIndex := (i+1)*blockSize - 1
		rightRes := querySingleElement(elements[middleIndex])
		rightIndex := middleIndex + 1 // now is the first index of the next block
		for block.Len() > 0 {
			q := block.Dequeue()
			for rightIndex < q.right { // Expand right side
				expandQ(&rightRes, elements[rightIndex])
				rightIndex++
			}
			res[q.index] = clone(rightRes) // Expand left side
			for q.left < middleIndex {
				expandQ(&res[q.index], elements[q.left])
				q.left++
			}
		}
	}
	return
}

func SelectionSort[T any](slice []T, Prior func(a, b T) bool) {
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if Prior(slice[j], slice[min]) {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}
}

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
