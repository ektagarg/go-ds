package sort

func SelectionSort(a []int) []int {
	// repeatedly find the minimum element and insert it in the beginning
	for i := 0; i < len(a); i++ {
		min := a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
			}
		}
		a[i], min = min, a[i]
	}
	return a
}
