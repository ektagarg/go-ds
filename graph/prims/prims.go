package graph

import (
	"fmt"
)

//======================================== Graph Structure Implementation starts========================================//

//Edge => an edge struct which represent a edge from vertex A to vertex B having X Weight
type Edge struct {
	From   int
	To     int
	Weight int
}

//Graph => graph data structure
type Graph struct {
	vertices int
	edges    [][]Edge
}

//Initialize => initializes the graph
func (g *Graph) Initialize(vertices int) {

	//set the vertices we have
	g.vertices = vertices

	// initialize the vertices
	if g.edges == nil {
		g.edges = make([][]Edge, vertices)
	}

	//initialize each vertex as  an empty array
	for i := 0; i < vertices; i++ {
		g.edges[i] = []Edge{}
	}
}

//AddEdge => weighted edge add
func (g *Graph) AddEdge(u int, v int, w int) {
	currEdge := Edge{u, v, w}

	//for undirected graph both should have the entry
	currEdgeU := Edge{v, u, w}

	//set the v is reachabe from u
	g.edges[u] = append(g.edges[u], currEdge)
	g.edges[v] = append(g.edges[v], currEdgeU)

}

//Print the graph to see the Adjacency List
func (g *Graph) Print() {

	for i := range g.edges {
		fmt.Printf("%v\n", g.edges[i])
	}

}

//========================================= Graph Implementation Ends Here =======================================//

func getMinimum(arr []Edge, visited []bool) Edge {

	min := 99999
	//default value
	minIdx := Edge{0, 0, 99999}
	for i := range arr {
		if arr[i].Weight < min && visited[arr[i].To] == false {

			minIdx = arr[i]
			min = arr[i].Weight
		}
	}
	return minIdx
}

/** PrimsAlgorithm
The prims Algoeithm selects the minimum spanning tree from a graph based on a greed approach the algorithm is as Follows
	1. Select The starting edge.
	2. Set the node as first node then choose the next minimum which is also connected to current selected nodes.
	3. repeat until all vertices are included
**/
func PrimsAlgorithm(g Graph, verbose bool) {

	var MST []Edge
	visited := make([]bool, g.vertices)
	totalEdges := g.vertices - 1

	//set the 0 as visited node
	visited[0] = true

	//In end we want total vertex - 1 edges to connect all the vertex so running a for loop for that
	for i := 0; i < totalEdges; i++ {

		//Now we will select the minimum edge among all the edges that are connected to current selected vertices based on greedy approach
		//for that we will check all the edges that are connected to visited node.
		//then we will choose the minimum among them
		fmt.Printf("\tCurrently visiting edge  %v \n", i)

		//this variable will hold the value of  the minimum edge in the current iteration
		var curEdges []Edge
		for idx := range visited {

			//we explore the edges of only those vertex tha are already visited
			if visited[idx] == true {

				//get the minimum weight edge among the edges that originated from the current vertex
				minEdge := getMinimum(g.edges[idx], visited)

				//append this in the current set of minimum edges
				curEdges = append(curEdges, minEdge)
			}
		}

		//out of all the minimum edges from all the vertices choose one which is minimum among them
		minEdge := getMinimum(curEdges, visited)
		fmt.Printf("\tFound MinEdge %v  \n\n\n ", minEdge)

		//append that to the Minimum spanning tree
		MST = append(MST, minEdge)
		visited[minEdge.To] = true

	}
	fmt.Println("PRINTING THE MST")
	fmt.Println(MST)

}
