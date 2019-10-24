package sort

// BitonicSort sorts slice using parallel sorting algorithm,
// which performs O(n2 log n) comparisons.
//
// This sort works when size of input is power of 2.
//
func BitonicSort(array []int) {
	bitonicSort(array, 0, len(array), true)
}

func bitonicSort(array []int, low, count int, d bool) {
	if count < 2 {
		return
	}
	k := count / 2
	bitonicSort(array, low, k, true)
	bitonicSort(array, low+k, k, false)
	merge(array, low, count, d)
}

func merge(array []int, low, count int, d bool) {
	if count < 2 {
		return
	}
	k := count / 2
	for i := 0; i < low+k; i++ {
		if d != (array[i] > array[i+k]) {
			continue
		}
		array[i], array[i+k] = array[i+k], array[i]
	}
	merge(array, low, k, d)
	merge(array, low+k, k, d)
}
