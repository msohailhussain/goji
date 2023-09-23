// dir graph, undir graph, weighted, unweighted
package graph

import (
	cl "github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils"
	constr "github.com/lorenzotinfena/goji/utils/constraints"
)

type ShortedPathVertex[V comparable, W constr.Integer | constr.Float] struct {
	Vertex   V
	Previous *ShortedPathVertex[V, W]
	Cost     W
}

type unitGraphDijkstraIterator[V comparable] struct {
	g       UnitGraph[V]
	toVisit cl.Queue[ShortedPathVertex[V, int]]
	visited cl.Set[V]
}

func (it *unitGraphDijkstraIterator[V]) HasNext() bool {
	return it.toVisit.Len() != 0
}

func (it *unitGraphDijkstraIterator[V]) Next() ShortedPathVertex[V, int] {
	cur := it.toVisit.Dequeue()
	for _, v := range it.g.getAdjacents(cur.Vertex) {
		if !it.visited.Contains(v) {
			it.toVisit.Enqueue(ShortedPathVertex[V, int]{Vertex: v, Previous: &cur, Cost: cur.Cost + 1})
		}
	}
	return cur
}

func (g UnitGraph[V]) Dijkstra(from V) utils.Iterator[ShortedPathVertex[V, int]] {
	toVisit := *cl.NewQueue[ShortedPathVertex[V, int]](nil)
	toVisit.Enqueue(ShortedPathVertex[V, int]{Vertex: from, Previous: nil, Cost: 0})

	return &unitGraphDijkstraIterator[V]{
		toVisit: toVisit,
		visited: *cl.NewSet[V](),
	}
}

/* //TODO
type weightedGraphDijkstraIterator[V comparable] struct {
	bfsIterator utils.Iterator[V]
	last        *ShortedPathVertex[V, int]
}

func (it *unitGraphDijkstraIterator[V]) HasNext() bool {
	return it.bfsIterator.HasNext()
}

func (it *unitGraphDijkstraIterator[V]) Next() ShortedPathVertex[V, int] {
	current := it.bfsIterator.Next()
	if it.last == nil {
		return ShortedPathVertex[V, int]{
			Vertex:   current,
			Previous: it.last,
			Cost:     0,
		}
	} else {
		return ShortedPathVertex[V, int]{
			Vertex:   current,
			Previous: it.last,
			Cost:     it.last.Cost + 1,
		}
	}
}

func (g UnitGraph[V]) Dijkstra(from V) utils.Iterator[ShortedPathVertex[V, int]] {
	return &unitGraphDijkstraIterator[V]{
		bfsIterator: g.GetIteratorBFS(from),
		last:        nil,
	}
}
*/
