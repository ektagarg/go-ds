package sort

// CycleSort ...
// 1) get element
// 2) Find that element position
// 3) place it to right place
// 4) get previous position element, go to step 2
func CycleSort(a []int) []int {
	for cycleStart, mainCycleItem := range a {
		item := mainCycleItem
		pos := cycleStart
		pos = getElementRealPosition(pos, cycleStart, item, a)
		if pos == cycleStart {
			continue // element at right place!
		}
		for item == a[pos] {
			pos += 1
		}
		a[pos], item = item, a[pos]
		for pos != cycleStart { // pos == cycleStart => new element at right place!
			pos = getElementRealPosition(pos, cycleStart, item, a)
			for item == a[pos] {
				pos += 1
			}
			a[pos], item = item, a[pos]
		}
	}
	return a
}

func getElementRealPosition(pos, cycleStart, item int, a []int) int {
	pos = cycleStart
	for i := cycleStart + 1; i < len(a); i++ {
		if a[i] < item {
			pos += 1
		}
	}
	return pos
}
