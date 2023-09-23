// dir graph, undir graph, weighted, unweighted
package graph

import "github.com/lorenzotinfena/goji/utils/constraints"

func toString[V comparable, W constraints.Integer | constraints.Float, G UnitGraph[V] | WeightedGraph[V, W]](g G) string {
	return ""
}

func (g UnitGraph[V]) String() string {
	return toString[V, int](g)
}

func (g WeightedGraph[V, W]) String() string {
	return toString[V, W](g)
}
