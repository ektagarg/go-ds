package sort

//BubbleSort ... It sorts by comparing and sorting adjacent elments
func BubbleSort(a []int) []int {
	swapped := false //To check if list is already sorted; then return;

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				//swap elements
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		if !swapped {
			return a
		}
	}
	return a
}
