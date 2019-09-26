package sort

func partition(a []int, low int, high int) int {
	pivot := a[high]
	partitionIndex := low
	for i := low; i < high; i++ {
		if a[i] <= pivot {
			a[i], a[partitionIndex] = a[partitionIndex], a[i]
			partitionIndex++
		}
	}
	a[high], a[partitionIndex] = a[partitionIndex], a[high]
	return partitionIndex

}

// Another partitioning technique
func partitionTwo(a []int, low int, high int) int {
	pivotIndex := low - 1
	pivot := a[low-1]
	swapped := false

	for low < high {
		// stop if high is less than pivot
		for a[high] > pivot && low < high {
			high--

		}
		//stop is low greater than pivot
		for a[low] < pivot && low < high {
			low++

		}
		//swap high and low
		a[low], a[high] = a[high], a[low]

		if low == high && a[low] < a[pivotIndex] {
			a[pivotIndex], a[low] = a[low], a[pivotIndex]
			swapped = true
		}
	}
	if !swapped {
		return low
	}
	return pivotIndex

}

func QuickSort(a []int, low int, high int) {

	if low < high {
		pivotPosition := partition(a, low, high)

		QuickSort(a, low, pivotPosition-1)

		QuickSort(a, pivotPosition+1, high)
	}
}
