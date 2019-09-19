// Package search is an implementation of various search algorithm
package search

// LinearSearch ... Traversing the array to find the target element by comparing it with every other element
func LinearSearch(a []int, element int) int {
	for i := 0; i < len(a); i++ {
		if element == a[i] {
			return i + 1
		}
	}
	return 0
}
