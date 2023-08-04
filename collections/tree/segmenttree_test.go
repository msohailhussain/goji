package tree_test

import (
	"fmt"
	"testing"

	"github.com/lorenzotinfena/goji/collections/tree"
)

func TestBitset(t *testing.T) {
	s := tree.NewSegmentTree[int, int](
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		func(element int) int { return element },
		func(q1, q2 int) int { return q1 + q2 },
		func(oldQ, oldE, newE int) (newQ int) { return oldQ - oldE + newE },
	)

	fmt.Println(s)
}
