package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestMergeSort ... Test if the program actually sorts array or not
func TestMergeSort(t *testing.T) {
	a := []int{12, 90, 23, 87, 32, 67, 27}
	low := 0
	high := len(a) - 1
	MergeSort(a, low, high)
	result := []int{12, 23, 27, 32, 67, 87, 90}

	assert.Equal(t, result, a, "Oops.. Merge sort didn't work!!")
}
