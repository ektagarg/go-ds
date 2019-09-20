package main

import (
	"fmt"
	"log"
	"time"

	bs "./Arrays/Search/binary_search"
	ls "./Arrays/Search/linear_search"
	bbs "./Arrays/Sort/bubble_sort"
	is "./Arrays/Sort/insertion_sort"
	ss "./Arrays/Sort/selection_sort"
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
