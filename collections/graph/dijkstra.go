// dir graph, undir graph, weighted, unweighted
package graph

import (
	"github.com/lorenzotinfena/goji/utils"
	constr "github.com/lorenzotinfena/goji/utils/constraints"
)

type ShortedPathVertex[V comparable, W constr.Integer | constr.Float] struct {
	Vertex   V
	Previous *ShortedPathVertex[V, W]
	Cost     W
}

type unitGraphDijkstraIterator[V comparable] struct {
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
