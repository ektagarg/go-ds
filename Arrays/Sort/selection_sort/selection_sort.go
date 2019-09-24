package sort

//SelectionSort ... repeatedly find the minimum element and insert it in the beginning
func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := a[i]
		minLoc := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				minLoc = j
			}
		}
		a[i], a[minLoc] = a[minLoc], a[i]
	}
	return a
}
