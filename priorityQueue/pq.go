//Package heap provides heap operations for any type that implements heap.Interface.
//A heap is a tree with the property that each node is the minimum-valued node in its subtree.

package priorityqueue

import "container/heap"

//Item ... priority queue item
type Item struct {
	Data     string
	Priority int
	Index    int
}

//PriorityQueue .. priority queue implements heap.interface
type PriorityQueue []*Item

// return length of priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// returns the highest priority element
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

// swap elements
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push .. push items one by one to the prority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := pq.Len()
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop .. pop last item from the prority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old.Len()
	toRemove := old[n-1]

	// avoid memeory leak
	old[n-1] = nil
	toRemove.Index = -1

	// update the priority queue
	*pq = old[0 : n-1]
	return toRemove
}

// Update .. Update the item in the heap
func (pq *PriorityQueue) Update(item *Item, data string, priority int) {
	item.Data = data
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
