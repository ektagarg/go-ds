package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCocktailSort(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		valid []int
	}{
		{
			"Unsorted array",
			[]int{5, 3, 2, 1, 8, 4},
			[]int{1, 2, 3, 4, 5, 8},
		},
		{
			"Sorted array",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"Equal values",
			[]int{5, 5, 5, 5, 5, 5},
			[]int{5, 5, 5, 5, 5, 5},
		},
		{
			"Sorted and reversed",
			[]int{6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.valid, CocktailSort(test.arr))
		})
	}
}
