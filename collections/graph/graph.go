// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/lorenzotinfena/goji/utils/constraints"
	"github.com/lorenzotinfena/goji/utils/slices"
)

type unitGraph[V constraints.Equalized] struct {
	Edges map[V]cl.Set[V]
}

type weightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float] struct {
	Edges map[V]cl.Set[cl.Pair[V, W]]
}

func NewUnitGraph[V constraints.Equalized]() *unitGraph[V] {
	return &unitGraph[V]{
		Edges: make(map[V]cl.Set[V]),
	}
}

func NewWeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float]() *weightedGraph[V, W] {
	return &weightedGraph[V, W]{
		Edges: make(map[V]cl.Set[cl.Pair[V, W]]),
	}
}

func (g *unitGraph[V]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = *cl.NewSet[V]()
	}
}

func (g *weightedGraph[V, W]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = *cl.NewSet[cl.Pair[V, W]]()
	}
}

// Assumptions:
// - source vertex exist
func (g *unitGraph[V]) AddEdge(source, dest V) {
	e := g.Edges[source]
	e.Add(dest)
}

// Assumptions:
// - source vertex exist
func (g *weightedGraph[V, W]) AddEdge(source, dest V, weight W) {
	e := g.Edges[source]
	e.Add(cl.MakePair(dest, weight))
}

// Assumptions:
// - source vertex exist
func (g unitGraph[V]) getAdjacents(v V) []V {
	return g.Edges[v].ToSlice()
}

// Assumptions:
// - source vertex exist
func (g weightedGraph[V, W]) getAdjacents(v V) []V {
	return slices.Map(g.Edges[v].ToSlice(), func(p cl.Pair[V, W]) V { return p.First })
}

func (g *unitGraph[V]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS[V](root, func(v V) []V { return g.getAdjacents(v) })
}

func (g *weightedGraph[V, _]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS[V](root, func(v V) []V { return g.getAdjacents(v) })
}

func (g *unitGraph[V]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS[V](root, func(v V) []V { return g.getAdjacents(v) })
}

func (g *weightedGraph[V, _]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS[V](root, func(v V) []V { return g.getAdjacents(v) })
}
