package tree

import "github.com/lorenzotinfena/goji/collections/graph"

type TreeNode[V comparable] struct {
	Value    V
	Children []*TreeNode[V]
}

func (t *TreeNode[V]) ToGraph() graph.UnitGraph[V] {
	g := *graph.NewUnitGraph[V]()
	var build func(t *TreeNode[V])
	build = func(t *TreeNode[V]) {
		for _, c := range t.Children {
			g.AddVertex(c.Value)
			g.AddEdge(t.Value, c.Value)
			build(c)
		}
	}
	g.AddVertex(t.Value)
	build(t)
	return g
}

func (t *TreeNode[V]) String() string {
	return t.ToGraph().String()
}
