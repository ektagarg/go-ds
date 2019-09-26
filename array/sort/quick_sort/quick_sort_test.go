package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestQuickSort ... Test if the program actually sorts array or not
func TestQuickSort(t *testing.T) {
	a := []int{12, 90, 23, 87, 32, 67, 27}
	low := 1
	high := len(a) - 1
	QuickSort(a, low, high)
	result := []int{12, 23, 27, 32, 67, 87, 90}

	assert.Equal(t, result, a, "Oops.. Quick sort didn't work!!")
}
