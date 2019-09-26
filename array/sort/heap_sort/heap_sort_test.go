package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestHeapSort ... Test if the program actually sorts the given array or not
func TestHeapSort(t *testing.T) {
	a := []int{12, 90, 23, 87, 32, 67, 27}
	n := len(a)
	HeapSort(a, n)
	result := []int{12, 23, 27, 32, 67, 87, 90}

	assert.Equal(t, result, a, "Oops.. Heap sort didn't work!!")
}
