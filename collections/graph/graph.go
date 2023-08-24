// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

type unitGraph[V constraints.Equalized] struct {
	Edges map[V]cl.Set[V]
}

type weightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Integer] struct {
	Edges map[V]cl.Set[cl.Pair[V, W]]
}

func NewUnitGraph[V constraints.Equalized]() *unitGraph[V] {
	return &unitGraph[V]{
		Edges: make(map[V]cl.Set[V]),
	}
}
func NewWeightedGraph[V constraints.Equalized, W constraints.Integer | constraints.Integer]() *weightedGraph[V, W] {
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
func (g *weightedGraph[V, W]) AddEdge(source, dest V, weight W) {
	e := g.Edges[source]
	e.Add(cl.MakePair(dest, weight))
}

// Assumptions:
// - source vertex exist
func (g *unitGraph[V]) AddEdge(source, dest V) {
	e := g.Edges[source]
	e.Add(dest)
}

// TODO:
/*
type dfsIterator[T comparable] struct {
	g       *Graph[T]
	toVisit collections.Stack[T]
	visited set.Set[T]
}

func (it *dfsIterator[T]) HasNext() bool {
	return it.toVisit.Len() != 0
}
func (it *dfsIterator[T]) Next() T {
	cur := it.toVisit.Pop()
	nexts := it.g.E[cur]
	for _, v := range nexts.ToSlice() {
		if !it.visited.Contains(v.nextVertex) {
			it.toVisit.Push(v.nextVertex)
		}
	}
	return cur
}
func (g *Graph[T]) IterateDFS(root T) utils.Iterator[T] {
	toVisit := *collections.NewStack[T](nil)
	toVisit.Push(root)

	return &dfsIterator[T]{
		g:       g,
		toVisit: toVisit,
		visited: *set.NewSet[T](),
	}
}

type bfsIterator[T comparable] struct {
	g       *Graph[T]
	toVisit collections.Queue[T]
	visited set.Set[T]
}

func (it *bfsIterator[T]) HasNext() bool {
	return it.toVisit.Len() != 0
}
func (it *bfsIterator[T]) Next() T {
	cur := it.toVisit.Dequeue()
	nexts := it.g.E[cur]
	for _, v := range nexts.ToSlice() {
		if !it.visited.Contains(v.nextVertex) {
			it.toVisit.Enqueue(v.nextVertex)
		}
	}
	return cur
}
func (g *Graph[T]) IterateBFS(root T) utils.Iterator[T] {
	toVisit := *collections.NewQueue[T](nil)
	toVisit.Enqueue(root)

	return &bfsIterator[T]{
		g:       g,
		toVisit: toVisit,
		visited: *set.NewSet[T](),
	}
}*/
