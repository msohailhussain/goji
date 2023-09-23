// dir graph, undir graph, weighted, unweighted
package graph_test

import (
	"fmt"
	"testing"

	"github.com/lorenzotinfena/goji/collections/graph"
)

func TestGraphDrawing(t *testing.T) {
	g := graph.NewUnitGraph[int]()
	for i := 0; i < 10; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(0, 5)
	g.AddEdge(1, 6)
	g.AddEdge(1, 8)
	g.AddEdge(8, 9)
	g.AddEdge(9, 8)
	g.AddEdge(3, 7)
	g.AddEdge(7, 3)
	g.AddEdge(5, 4)
	g.AddEdge(7, 6)
	g.AddEdge(1, 9)
	g.AddEdge(2, 1)
	g.AddEdge(3, 2)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)
	g.AddEdge(7, 2)
	g.AddEdge(5, 5)
	g.AddEdge(9, 9)
	fmt.Println(g)
}
