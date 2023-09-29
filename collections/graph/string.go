// dir graph, undir graph, weighted, unweighted
package graph

import (
	"encoding/json"
	"fmt"

	"github.com/lorenzotinfena/goji/utils/constraints"
)

func toString[V comparable, W constraints.Integer | constraints.Float, G Graph[V, W]](g G) string {
	type NodeGraphData struct {
		ID    string `json:"id"`
		Label string `json:"label,omitempty"`
		Color string `json:"color,omitempty"`
		Shape string `json:"shape,omitempty"`
	}

	type EdgeGraphData struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Label  string `json:"label,omitempty"`
		ID     string `json:"id"`
		Color  string `json:"color,omitempty"`
		Dashes bool   `json:"dashes,omitempty"`
	}

	type Graph struct {
		Kind  map[string]bool `json:"kind"`
		Nodes []NodeGraphData `json:"nodes"`
		Edges []EdgeGraphData `json:"edges"`
	}

	graph := &Graph{
		Kind:  map[string]bool{"graph": true},
		Nodes: []NodeGraphData{},
		Edges: []EdgeGraphData{},
	}
	for _, v := range g.Vertices() {
		s := fmt.Sprint(v)
		graph.Nodes = append(graph.Nodes, NodeGraphData{ID: s, Label: s})
	}
	
	for _, v := range g.Vertices() {
		for _, adj := range g.GetAdjacents(v) {
			s1 := fmt.Sprint(v)
			s2 := fmt.Sprint(adj)
			tmp:= EdgeGraphData{From: s1, To: s2}
			gra, isWeighted := any(g).(WeightedGraph[V, W])
			if isWeighted {
				tmp.Label = fmt.Sprint(gra.GetWeight(v, adj))
			}
			graph.Edges = append(graph.Edges, tmp)
		}
		
	}
	s, _ := json.Marshal(graph)
	return string(s)
}

func (g UnitGraph[V]) String() string {
	return toString[V, int](g)
}

func (g WeightedGraph[V, W]) String() string {
	return toString[V, W](g)
}
