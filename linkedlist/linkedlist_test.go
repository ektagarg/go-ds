package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning

var linkedlist = NewLinkedList()

func TestInsertBeginning(t *testing.T) {
	node1 := NewNode(2)
	node2 := NewNode(34)
	node3 := NewNode(93)
	linkedlist.InsertBeginning(node1)
	linkedlist.InsertBeginning(node2)
	linkedlist.InsertBeginning(node3)
	assert.Equal(t, 3, linkedlist.Len, "InsertBeginning: Linked list length is incorrect!")

	// check if the first element is actually 93
	head := linkedlist.Start
	assert.Equal(t, 93, head.Data, "Linkedlist InsertBeginning: didn't work!")
}

func TestInsertEnd(t *testing.T) {
	// insert in end
	node4 := NewNode(222)
	linkedlist.InsertEnd(node4)
	assert.Equal(t, 4, linkedlist.Len, "InsertEnd: Linked list length is incorrect!")

	// check if the last element is actually 222
	end := linkedlist.End
	assert.Equal(t, 222, end.Data, "Linkedlist InsertEnd: didn't work!")
}

func TestInsertAtPosition(t *testing.T) {
	//insert at a particular position
	node5 := NewNode(555)
	linkedlist.InsertAtPosition(node5, 3)
	assert.Equal(t, 5, linkedlist.Len, "InsertAtPosition: Linked list length is incorrect!")

	count := 1
	head := linkedlist.Start

	for {
		if count == 3 {
			assert.Equal(t, 555, head.Data, "Linkedlist InsertAtPosition: didn't work!")
			break
		}
		head = head.Next
		count++
	}
}

func TestDeleteBeginning(t *testing.T) {
	// delete from beginnning
	linkedlist.DeleteBeginning()
	assert.Equal(t, 4, linkedlist.Len, "DeleteBeginning: Linked list length is incorrect!")
}

func TestDeleteEnd(t *testing.T) {
	// delete from end
	linkedlist.DeleteEnd()
	assert.Equal(t, 3, linkedlist.Len, "DeleteEnd: Linked list length is incorrect!")
}

func TestDeleteFromPosition(t *testing.T) {
	// delete from any position
	linkedlist.DeleteFromPosition(2)
	assert.Equal(t, 2, linkedlist.Len, "DeleteFromPosition: Linked list length is incorrect!")
}
