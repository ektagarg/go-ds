package sort

func MergeSort(a []int, l int, r int) {
	if l < r {
		// get the middle element
		m := (l + r) / 2

		//divide slice from l to m
		MergeSort(a, l, m)
		//divide slice from m+1 to r
		MergeSort(a, m+1, r)

		//merge slices
		merge(a, l, m, r)
	}

}

func merge(a []int, low int, mid int, high int) {
	// create a temp slice to insert the elements i sorted order
	temp := make([]int, high-low+1)
	i := low
	j := mid + 1
	k := 0

	// traverse both slices
	for i <= mid && j <= high {
		if a[i] < a[j] {
			// smaller element at the right side; insert it to temp array
			temp[k] = a[i]
			i++
		} else {
			// smaller element at the left side; insert it to temp array
			temp[k] = a[j]
			j++
		}
		k++
	}

	// copy rest of the elements of slice - left to mid into temp array
	for i <= mid {
		temp[k] = a[i]
		k++
		i++
	}

	// copy rest of the elements of slice - mid+1 to right into temp array
	for j <= high {
		temp[k] = a[j]
		k++
		j++
	}

	// add elements to the given array
	for i := low; i <= high; i++ {
		a[i] = temp[i-low]
	}
}
