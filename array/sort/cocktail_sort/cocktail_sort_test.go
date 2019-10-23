package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//TestCocktailSort ... Test if the program actually sorts array or not
func TestCocktailSort(t *testing.T) {
	a := [][]int{
		{5, 1, 4, 2, 8, 0, 2},
		{0, -1, 2, 2, -4, 5, -8},
		{12, 90, 23, 87, 32, 67, 27},
	}
	result := [][]int{
		{0, 1, 2, 2, 4, 5, 8},
		{-8, -4, -1, 0, 2, 2, 5},
		{12, 23, 27, 32, 67, 87, 90},
	}
	for i := range a {
		sorted := CocktailSort(a[i])

		assert.Equal(t, result[i], sorted, "Oops.. cocktail sort didn't work!!")

	}

}
