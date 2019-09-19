package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestBinarySearch ... Check if the program actually finds the correct position of element in the sorted array
func TestBinarySearch(t *testing.T) {
	a := []int{12, 23, 27, 32, 67, 87, 90}
	position := BinarySearch(a, 67)

	assert.Equal(t, 5, position, "Oops.. binary search didn't work!!")
}
