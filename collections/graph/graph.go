// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

type Graph[V comparable, W constraints.Integer | constraints.Float] interface {
	UnitGraph[V] | WeightedGraph[V, W]
	GetAdjacents(V) []V
	GetIteratorDFS(root V) utils.Iterator[V]
	GetIteratorBFS(root V) utils.Iterator[V]
	Vertices() []V
	String() string
}

// Every vertex is unique
type UnitGraph[V constraints.Equalized] struct {
	Edges map[V]cl.HashSet[V]
}

// Every vertex is unique
type WeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float] struct {
	Edges map[V]map[V]W
}

func NewUnitGraph[V constraints.Equalized]() *UnitGraph[V] {
	return &UnitGraph[V]{
		Edges: make(map[V]cl.HashSet[V]),
	}
}

func NewWeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Float]() *WeightedGraph[V, W] {
	return &WeightedGraph[V, W]{
		Edges: make(map[V]map[V]W),
	}
}

func (g *UnitGraph[V]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = *cl.NewHashSet[V]()
	}
}

func (g *WeightedGraph[V, W]) AddVertex(v V) {
	_, present := g.Edges[v]
	if !present {
		g.Edges[v] = make(map[V]W)
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
	e[dest] = weight
}

// Assumptions:
// - source vertex exist
func (g UnitGraph[V]) GetAdjacents(v V) []V {
	return g.Edges[v].ToSlice()
}

// Assumptions:
// - source vertex exist
func (g WeightedGraph[V, W]) GetAdjacents(v V) []V {
	keys := make([]V, 0, len(g.Edges[v]))
	for k := range g.Edges[v] {
		keys = append(keys, k)
	}
	return keys
}

func (g UnitGraph[V]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS(root, func(v V) []V { return g.GetAdjacents(v) })
}

func (g WeightedGraph[V, _]) GetIteratorDFS(root V) utils.Iterator[V] {
	return cl.NewIteratorDFS(root, func(v V) []V { return g.GetAdjacents(v) })
}

func (g UnitGraph[V]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS(root, func(v V) []V { return g.GetAdjacents(v) })
}

func (g WeightedGraph[V, _]) GetIteratorBFS(root V) utils.Iterator[V] {
	return cl.NewIteratorBFS(root, func(v V) []V { return g.GetAdjacents(v) })
}

func (g UnitGraph[V]) Vertices() []V {
	keys := make([]V, len(g.Edges))
	i := 0
	for k := range g.Edges {
		keys[i] = k
		i++
	}
	return keys
}

func (g WeightedGraph[V, W]) Vertices() []V {
	keys := make([]V, len(g.Edges))
	i := 0
	for k := range g.Edges {
		keys[i] = k
		i++
	}
	return keys
}

func (g WeightedGraph[V, W]) GetWeight(from, to V) W {
	return g.Edges[from][to]
}
