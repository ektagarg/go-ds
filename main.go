package main

import (
	"fmt"
	"log"
	"time"

	bs "./array/search/binary_search"
	ls "./array/search/linear_search"
	bbs "./array/sort/bubble_sort"
	is "./array/sort/insertion_sort"
	ss "./array/sort/selection_sort"
)

func main() {
	a := []int{23, 45, 67, 27, 32, 87, 90, 12}
	position := ls.LinearSearch(a, 67)
	fmt.Println(position)

	b := []int{12, 23, 27, 32, 67, 87, 90}
	positions := bs.BinarySearch(b, 67)
	fmt.Println(positions)

	start := time.Now()
	bs := bbs.BubbleSort(a)
	elapsed := time.Since(start)
	fmt.Println(bs)
	log.Printf("bubble took %s", elapsed)

	start = time.Now()
	is := is.InsertionSort(a)
	elapsed = time.Since(start)
	fmt.Println(is)
	log.Printf("Insertion took %s", elapsed)

	start = time.Now()
	ss := ss.SelectionSort(a)
	elapsed = time.Since(start)
	fmt.Println(ss)
	log.Printf("Selection took %s", elapsed)
}
