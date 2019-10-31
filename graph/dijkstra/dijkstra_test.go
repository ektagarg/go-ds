package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := [][]int32{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0}}
	distances := []int32{0, 4, 12, 19, 21, 11, 9, 8, 14}

	assert.Equal(t, Dijkstra(graph, 0), distances, "Oops.. Dijkstra shortest distances algorithm didn't work!!")
}
