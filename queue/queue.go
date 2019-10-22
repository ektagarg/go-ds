package queue

import "fmt"

//Queue ... queue data structure
type Queue struct {
	Data []int
}

//NewQueue ... create an empty queue
func NewQueue() Queue {
	return Queue{
		Data: make([]int, 0),
	}
}

//IsEmpty ... check if queue is empty
func (q *Queue) IsEmpty() bool {
	if len(q.Data) == 0 {
		return true
	}
	return false
}

//Size ... returns size of queue
func (q *Queue) Size() int {
	return len(q.Data)
}

//Front ... first element from the queue
func (q *Queue) Front() int {
	if !(q.IsEmpty()) {
		return q.Data[0]
	}
	return -1
}

//Enqueue ... add a element to queue
func (q *Queue) Enqueue(element int) {
	// if queue is empty
	if len(q.Data) == 0 {
		q.Data = append(q.Data, element)
		return
	}
	i := 1
	for {
		if i == len(q.Data) {
			q.Data = append(q.Data, element)
			return
		}
		i++
	}
}

// Dequeue ... delete an element from the queue
func (q *Queue) Dequeue() int {
	deleted := q.Front()
	q.Data = q.Data[1:]
	return deleted
}

// PrintQueue ... print the queue
func (q *Queue) PrintQueue() {
	for i := 0; i < len(q.Data); i++ {
		fmt.Println(q.Data[i])
	}
}
