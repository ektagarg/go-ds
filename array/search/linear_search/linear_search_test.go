package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestLinearSearch ... Check if the program actually finds the correct position of element in the array
func TestLinearSearch(t *testing.T) {
	a := []int{23, 45, 67, 27, 32, 87, 90, 12}
	position := LinearSearch(a, 67)

	assert.Equal(t, 3, position, "Oops.. linear search didn't work!!")
}
