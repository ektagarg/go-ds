package doublylinklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning

var dll = NewDoublylinklist()

func TestInsertBeginning(t *testing.T) {
	node1 := NewNode(2)
	node2 := NewNode(34)
	node3 := NewNode(93)
	dll.InsertBeginning(node1)
	dll.InsertBeginning(node2)
	dll.InsertBeginning(node3)
	assert.Equal(t, 3, dll.Length, "InsertBeginning: Doubly Linked list length is incorrect!")

	// check if the first element is actually 93
	head := dll.Start
	assert.Equal(t, 93, head.Data, "dll InsertBeginning: didn't work!")
}

func TestInsertEnd(t *testing.T) {
	// insert in end
	node4 := NewNode(222)
	dll.InsertEnd(node4)
	assert.Equal(t, 4, dll.Length, "InsertEnd: Doubly Linked list length is incorrect!")

	// check if the last element is actually 222
	end := dll.End
	assert.Equal(t, 222, end.Data, "dll InsertEnd: didn't work!")
}

func TestInsertMiddle(t *testing.T) {
	//insert at a particular position
	node5 := NewNode(555)
	dll.InsertMiddle(node5, 3)
	assert.Equal(t, 5, dll.Length, "InsertMiddle: Doubly Linked list length is incorrect!")

	count := 1
	head := dll.Start

	for {
		if count == 3 {
			assert.Equal(t, 555, head.Data, "dll InsertAtPosition: didn't work!")
			break
		}
		head = head.Next
		count++
	}
}
