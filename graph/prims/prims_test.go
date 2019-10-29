package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	var graph = Graph{}
	graph.Initialize(6)
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 3, 8)
	graph.AddEdge(3, 4, 1)
	graph.AddEdge(3, 5, 7)
	graph.AddEdge(4, 5, 6)
	graph.AddEdge(2, 5, 2)
	graph.AddEdge(1, 3, 11)
	graph.AddEdge(1, 2, 8)
	graph.Print()
	PrimsAlgorithm(graph)

}
