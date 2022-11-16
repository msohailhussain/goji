package main

import (
	"bufio"
	"fmt"
	"os"
)

// For LeetCode (copy paste only your solve function and all code below it)
func main() {
	io := newStdIO()
	defer io.flush()
	io.printLn( /* CALL SOLVE FUNCTION HERE */ )
}
// YOUR SOLVE FUNCTION HERE

// For GoogleKickStart
func main() {
	io := newStdIO()
	defer io.flush()
	T := io.ScanUInt16()
	for t := uint16(1); t <= T; t++ {
		io.print("Case #", t, ": ", solve(&io))
	}
}
func solve(io *IO) string {
	// SOLVE HERE
	return fmt.Sprintln( /* SOLUTIONS HERE */ )
}

// For Hackerrank
func main() {
	io := newStdIO()
	defer io.flush()
	for T := io.ScanUInt16(); T > 0; T-- {
		io.print(solve(&io))
	}
}
func solve(io *IO) string {
	// SOLVE HERE
	return fmt.Sprintln( /* SOLUTIONS HERE */ )
}

// For Codeforces
func main() {
	io := newFileIO()
	defer io.flush()
	for T := io.ScanUInt16(); T > 0; T-- {
		io.print(solve(&io))
	}
}
func solve(io *IO) string {
	// SOLVE HERE
	return fmt.Sprintln( /* SOLUTIONS HERE */ )
}

//#region INTERFACES

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

//#endregion

// #region TYPES AND METHODS
// #region SingleLinkedList
type singleLinkedListNode[T comparable] struct {
	Value T
	Next  *singleLinkedListNode[T]
}

type SingleLinkedList[ValueType comparable, IndexType Unsigned] struct {
	first  *singleLinkedListNode[ValueType]
	last   *singleLinkedListNode[ValueType]
	length IndexType
}

func NewSingleLinkedList[T comparable, I Unsigned]() *SingleLinkedList[T, I] {
	return &SingleLinkedList[T, I]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}
func (l *SingleLinkedList[T, I]) GetLength() I { return l.length }

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
		if tmp.Value == value {	return true }
		tmp = tmp.Next
	}
	return false
}
func (l *SingleLinkedList[T, I]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}
func (l *SingleLinkedList[T, I]) ToString() string {
	slice := make([]any, l.length)
	tmp := l.first
	for i := I(0); i < l.length; i++ {
		slice[i] = tmp.Value
		tmp = tmp.Next
	}
	return fmt.Sprint(slice)
}

//#endregion
//#endregion

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

func First[T any](a T, _ any) T { return a }

//#endregion

// #region IO STUFF
type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func newStdIO() IO {
	return IO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

// Is assumed that input.txt file exists
func newFileIO() IO {
	return IO{
		r: bufio.NewReader(First(os.Open("input.txt"))),
		w: bufio.NewWriter(First(os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm))),
	}
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

func (io *IO) print(x ...any)   { fmt.Fprint(io.w, x...) }
func (io *IO) printLn(x ...any) { fmt.Fprintln(io.w, x...) }

func (io *IO) flush() { io.w.Flush() }

//#endregion
