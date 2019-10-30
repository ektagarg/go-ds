package graph

import (
	"fmt"
	"sort"
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

/*
# A utility function to find set of an element i


    # A function that does union of two sets of x and y
    # (uses union by rank)
    def union(self, parent, rank, x, y):
        xroot = self.find(parent, x)
        yroot = self.find(parent, y)

        # Attach smaller rank tree under root of
        # high rank tree (Union by Rank)
        if rank[xroot] < rank[yroot]:
            parent[xroot] = yroot
        elif rank[xroot] > rank[yroot]:
            parent[yroot] = xroot

        # If ranks are same, then make one as root
        # and increment its rank by one
        else :
            parent[yroot] = xroot
            rank[xroot] += 1
*/

//traverses the parents array and fetches the representative parent of the vertex x
func findParent(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return findParent(parent, parent[i])
}

//Merges two vertices and sets the parent by rank
func union(parent []int, rank []int, x int, y int) {
	xroot := findParent(parent, x)
	yroot := findParent(parent, y)

	if rank[x] > rank[y] {
		parent[yroot] = xroot
	} else if rank[x] < rank[y] {
		parent[xroot] = yroot
	} else {
		parent[yroot] = xroot
		rank[xroot] = 1
	}
}

// Kruskals algorithm
func Kruskals(g Graph) {

	//get all the edges in an array and then sort it
	sortedEdges := []Edge{}
	MST := []Edge{}

	//TODO : Use a min Heap to optimize the following
	for i := range g.edges {
		for edge := range g.edges[i] {
			sortedEdges = append(sortedEdges, g.edges[i][edge])
		}
	}

	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i].Weight < sortedEdges[j].Weight
	})

	//initialize parent and the rank
	parent := make([]int, g.vertices)
	rank := make([]int, g.vertices)

	for i := 0; i < g.vertices; i++ {
		parent[i] = i
		rank[i] = 0
	}

	i := 0
	e := 0

	//we need to select V-1 edges for all the edges to connect
	for e < g.vertices-1 {

		//select the minumum edge not choosen yet
		currEdge := sortedEdges[i]
		i++

		//find the parent of the vertices of current edge
		x := findParent(parent, currEdge.From)
		y := findParent(parent, currEdge.To)

		//if parent of x is not equal to parent of y means both have different sets and thus wont form a cycle when merged
		//so we select this edge and increase the edges count
		if x != y {
			e++
			//add it to mst
			MST = append(MST, currEdge)
			//since the vertices are now connected union both of them for future use.
			union(parent, rank, x, y)
		}

	}
	fmt.Println("PRINTING THE MST")
	fmt.Println(MST)

}
