package sort

import (
	"math"
)

// RadixSort ... A function to sort array using Radix sort algorithm
func RadixSort(a []int) []int {
	iterations := findLargest(a)
	SortedSlice := a
	for i := 0; i < iterations; i++ {
		// type casting base because Pow10 retuen float64
		base := int(math.Pow10(i))

		SortedSlice = sortIterationWise(SortedSlice, base)
	}
	return SortedSlice
}

//findLargest ... Find the largest number in the given array to get the iterations
func findLargest(a []int) int {
	max := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			max = a[i]
		} else {
			max = a[i+1]
		}
	}
	c := countDigits(max)
	return c
}

//countDigits ... This function count digit in an integer
func countDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

//sortIterationWise ... Sort the slice based on each digit: interation wise
func sortIterationWise(a []int, base int) []int {

	//initialize buckets
	m := make(map[int][]int)
	SortedSlice := []int{}

	// iteration one by one for each element of array:
	for j := 0; j < len(a); j++ {
		num := a[j]

		// get digit based on the base in each iteration
		digit := num / base % 10

		// insert value inside map which has digits as keys
		m[digit] = append(m[digit], num)

	}

	//appending map values to a slice
	for j := 0; j < len(a); j++ {
		SortedSlice = append(SortedSlice, m[j]...)
	}
	return SortedSlice
}
