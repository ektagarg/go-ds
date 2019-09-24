package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestShellSort ... Test if the program actually sorts array or not
func TestShellSort(t *testing.T) {
	a := []int{12, 90, 23, 87, 32, 67, 27}
	sorted := ShellSort(a)
	result := []int{12, 23, 27, 32, 67, 87, 90}

	assert.Equal(t, result, sorted, "Oops.. Shell sort didn't work!!")
}
