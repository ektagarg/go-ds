package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCycleSort ... Test if the program actually sorts array or not
func TestCycleSort(t *testing.T) {
	randomUnsortedSlice, randomSortedSlice := generateRandomSlices()
	assert.NotEqual(t, randomUnsortedSlice, randomSortedSlice)

	cycleSorted := CycleSort(randomUnsortedSlice)

	assert.Equal(t, randomSortedSlice, cycleSorted, "Oops.. Cycle sort didn't work!!")
}

// generateRandomSlices generate random slice and his sorted copy
func generateRandomSlices() (unsorted []int, sorted []int) {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(100)
	randomUnsortedSlice := make([]int, size, size)
	for i := 0; i < size; i++ {
		randomUnsortedSlice[i] = rand.Intn(999)
	}
	randomSortedSlice := append([]int(nil), randomUnsortedSlice...) // Notice the ... splat
	sort.Ints(randomSortedSlice)
	return randomUnsortedSlice, randomSortedSlice
}
