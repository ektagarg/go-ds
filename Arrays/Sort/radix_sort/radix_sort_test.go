package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestRadixSort ... Test if the program actually sorts array or not
func TestRadixSort(t *testing.T) {
	a := []int{213, 3, 45, 67, 78, 1, 34, 456, 44}
	sortedSlice := RadixSort(a)
	result := []int{1, 3, 34, 44, 45, 67, 78, 213, 456}

	assert.Equal(t, result, sortedSlice, "Oops.. Radix sort didn't work!!")
}
