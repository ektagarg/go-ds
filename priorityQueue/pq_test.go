package priorityqueue

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Data:     value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		Data:     "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	assert.Equal(t, 4, pq.Len(), "length of priority queue is incorrect!")

	pq.Update(item, item.Data, 5)

	// Take the items out; they arrive in decreasing priority order.
	removed := heap.Pop(&pq).(*Item)
	assert.Equal(t, "orange", removed.Data, "element data is incorrect")
	assert.Equal(t, 5, removed.Priority, "element priority is incorrect")
}
