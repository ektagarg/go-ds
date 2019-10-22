// Package search is an implementation of various search algorithm
package search

// BinarySearch ... BinarySearch is the most efficient searching algorithm having a run-time complexity of O(log2 N). This algorithm works only on a sorted list of elements.
func BinarySearch(a []int, target int) int {
	first := 0
	last := len(a) - 1

	for first <= last {
		mid := (first + last) / 2
		if a[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	if first == len(a) || a[first] != target {
		return 0
	}
	return first + 1

}
