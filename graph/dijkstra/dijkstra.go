package graph

//utility variable used only for comparison purposes
const maxInt32 int32 = 1<<31 - 1

func Dijkstra(graph [][]int32, sourceNode int32) []int32 {
	var distance = make([]int32, len(graph))
	var sptSet = make([]bool, len(graph))
	//initializing distance as unvisited
	for i := 0; i < len(graph); i++ {
		distance[i] = maxInt32
		sptSet[i] = false
	}
	//distance to the starting node itself is always zero
	distance[sourceNode] = 0

	for count := 0; count < len(graph)-1; count++ {
		u := minDistance(distance, sptSet)
		sptSet[u] = true

		for v := 0; v < len(graph); v++ {
			if !sptSet[v] && graph[u][v] != 0 && distance[u] != maxInt32 && distance[u]+graph[u][v] < distance[v] {
				distance[v] = distance[u] + graph[u][v]
			}
		}
	}

	return distance
}

func minDistance(dist []int32, sptSet []bool) int {
	var min = maxInt32
	var minIndex int
	for v := 0; v < len(dist); v++ {
		if !sptSet[v] && dist[v] < min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}
