package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	type Scalar interface {
		~int8 | ~uint8 | ~int16 | ~uint16 | ~int32  | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~uintptr | ~float32 | ~float64
	}

	func max[T Scalar](a T, b T) T {
		if a >= b { return a}
		return b
	}
*/
const MAXINT32 = 2147483647
const MININT32 = -2147483648

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve() {

}

func main() {
	// Leave only one:
	io := NewFileIO()
	//io := NewStdIO()

	io.Flush()
}

type IO struct {
	r *bufio.Reader
	w *bufio.Writer
}

func NewStdIO() IO {
	return IO{
		r: bufio.NewReader(os.Stdin),
		w: bufio.NewWriter(os.Stdout),
	}
}

// Is assumed that both files exists
func NewFileIO() IO {
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

func (io *IO) Print(x ...any)   { fmt.Fprint(io.w, x...) }
func (io *IO) PrintLn(x ...any) { fmt.Fprintln(io.w, x...) }

func (io *IO) Flush() {
	io.w.Flush()
}
