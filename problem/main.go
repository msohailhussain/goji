package main

import (
	"bufio"
	"fmt"
	"os"
)

// For LeetCode
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

////////////////////////////////////////////////////////////////////////////

////////////////
// INTERFACES //
////////////////

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

///////////
// TYPES //
///////////

// # TODO type SingleLinkedList

///////////////
// FUNCTIONS //
///////////////

func Max[T Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}
func Min[T Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

// ////////////
// IO STUFF //
// ////////////
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

// Is assumed that both files exists
func newFileIO() IO {
	in, _ := os.Open("input.txt")
	ou, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	return IO{
		r: bufio.NewReader(in),
		w: bufio.NewWriter(ou),
	}
}

func (io *IO) ScanInt8() (x int8)   { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt16() (x int16) { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt32() (x int32) { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanInt64() (x int64) { _, _ = fmt.Fscan(io.r, &x); return }

func (io *IO) ScanUInt8() (x uint8)   { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt16() (x uint16) { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt32() (x uint32) { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanUInt64() (x uint64) { _, _ = fmt.Fscan(io.r, &x); return }

func (io *IO) ScanFloat32() (x float32) { _, _ = fmt.Fscan(io.r, &x); return }
func (io *IO) ScanFloat64() (x float64) { _, _ = fmt.Fscan(io.r, &x); return }

func (io *IO) ScanString() (x string) { _, _ = fmt.Fscan(io.r, &x); return }

func (io *IO) print(x ...interface{})   { fmt.Fprint(io.w, x...) }
func (io *IO) printLn(x ...interface{}) { fmt.Fprintln(io.w, x...) }

func (io *IO) flush() { io.w.Flush() }
