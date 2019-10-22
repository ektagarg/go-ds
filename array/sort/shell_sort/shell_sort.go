package sort

// ShellSort ... Implementing algorithm based on insertion sort
func ShellSort(a []int) []int {
	Len := len(a)

	h := 1
	for h < Len/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		// h-sort the array
		for i := h; i < Len; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j = j - h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h = h / 3
	}
	return a
}
