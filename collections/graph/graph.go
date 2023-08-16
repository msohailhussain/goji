package graph

import (
	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/collections/set"
	"github.com/lorenzotinfena/goji/utils"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

type edge[T any] struct {
	nextVertex T
	weight     int
}

func makeEdge[T any](nextVertex T, weight int) edge[T] {
	return edge[T]{
		nextVertex: nextVertex,
		weight:     weight,
	}
}

// Weighted directed graph
type Graph[T constraints.Equalized] struct {
	V set.Set[T]
	E map[T]set.Set[edge[T]]
}

func NewGraph[T constraints.Equalized]() *Graph[T] {
	return &Graph[T]{
		V: *set.NewSet[T](),
		E: make(map[T]set.Set[edge[T]]),
	}
}

func (g *Graph[T]) AddVertex(item T) {
	g.V.Add(item)
	g.E[item] = *set.NewSet[edge[T]]()
}

func (g *Graph[T]) AddEdge(source, dest T, weight int) {
	e := g.E[source]
	e.Add(makeEdge(dest, weight))
}

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
}
