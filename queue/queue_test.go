package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning

var queue = NewQueue()

func TestEnQueue(t *testing.T) {
	queue.Enqueue(23)
	queue.Enqueue(244)
	queue.Enqueue(565)
	queue.Enqueue(54)
	queue.Enqueue(98)

	assert.Equal(t, 5, queue.Size(), "Queue size incorrect!")
}

func TestDequeue(t *testing.T) {
	element := queue.Dequeue()
	assert.Equal(t, 4, queue.Size(), "Queue size incorrect!")
	assert.Equal(t, 23, element, "Dequeued element is incorrect")
}
