package sort

import "fmt"

//CocktailSort ... Basically two sided bubble sort
func CocktailSort(list []int) []int {
	swapped := true
	start, end := 0, len(list)-1
	for swapped {
		swapped = false
		fmt.Println()

		for i := start; i < end; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
			fmt.Println(list)
		}
		fmt.Println()
		if !swapped {
			break
		}
		swapped = false
		end--
		for i := end - 1; i >= start; i-- {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
			fmt.Println(list)
		}

	}
	return list
}
