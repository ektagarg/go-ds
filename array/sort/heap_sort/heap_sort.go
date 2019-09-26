package sort

//HeapSort ... Function to do heap sort
func HeapSort(a []int, n int) {
	// create heap
	for i := n / 2; i >= 0; i-- {
		Heapify(a, n, i)
	}

	// one by one extract elements from heap
	for i := n - 1; i > 0; i-- {

		// move current root to end
		a[0], a[i] = a[i], a[0]

		// call heapify
		Heapify(a, i, 0)
	}
}

//Heapify ... Function to heapify a subtree
func Heapify(a []int, length int, root int) {
	// set root as the largest element
	max := root
	// set left
	left := 2*root + 1
	//set right
	right := 2*root + 2

	// check if left is greater than max and set as max
	if left < length && a[left] > a[max] {
		max = left
	}

	// check if right is greater than max and set as max
	if right < length && a[right] > a[max] {
		max = right
	}

	// if max is not on the root
	if root != max {
		//swap root with max element
		a[max], a[root] = a[root], a[max]

		// heapify the subtree
		Heapify(a, length, max)
	}
}
