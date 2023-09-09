// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/lorenzotinfena/goji/utils/constraints"
	"github.com/lorenzotinfena/goji/utils/slices"
)

type UnitGraph[V constraints.Equalized] struct {
	Edges map[V]cl.Set[V]
}

type WeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float] struct {
	Edges map[V]cl.Set[cl.Pair[V, W]]
}

func NewUnitGraph[V constraints.Equalized]() *UnitGraph[V] {
	return &UnitGraph[V]{
		Edges: make(map[V]cl.Set[V]),
	}
}

func NewWeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float]() *WeightedGraph[V, W] {
	return &WeightedGraph[V, W]{
		Edges: make(map[V]cl.Set[cl.Pair[V, W]]),
	}
}

func (g *UnitGraph[V]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = *cl.NewSet[V]()
	}
}

func (g *WeightedGraph[V, W]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = *cl.NewSet[cl.Pair[V, W]]()
	}
}

// Assumptions:
// - source vertex exist
func (g *UnitGraph[V]) AddEdge(source, dest V) {
	e := g.Edges[source]
	e.Add(dest)
}

// Assumptions:
// - source vertex exist
func (g *WeightedGraph[V, W]) AddEdge(source, dest V, weight W) {
	e := g.Edges[source]
	e.Add(cl.MakePair(dest, weight))
}

// Assumptions:
// - source vertex exist
func (g UnitGraph[V]) getAdjacents(v V) []V {
	return g.Edges[v].ToSlice()
}

// Assumptions:
// - source vertex exist
func (g WeightedGraph[V, W]) getAdjacents(v V) []V {
	return slices.Map(g.Edges[v].ToSlice(), func(p cl.Pair[V, W]) V { return p.First })
}

func (g UnitGraph[V]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS(root, func(v V) []V { return g.getAdjacents(v) })
}

func (g WeightedGraph[V, _]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS(root, func(v V) []V { return g.getAdjacents(v) })
}

func (g UnitGraph[V]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS(root, func(v V) []V { return g.getAdjacents(v) })
}

func (g WeightedGraph[V, _]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS(root, func(v V) []V { return g.getAdjacents(v) })
}
