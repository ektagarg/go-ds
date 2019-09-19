package main

import (
	"fmt"

	Search "./Arrays/Search"
)

func main() {
	a := []int{23, 45, 67, 27, 32, 87, 90, 12}
	position := Search.LinearSearch(a, 67)
	fmt.Println(position)

	b := []int{12, 23, 27, 32, 67, 87, 90}
	positions := Search.BinarySearch(b, 67)
	fmt.Println(positions)
}
