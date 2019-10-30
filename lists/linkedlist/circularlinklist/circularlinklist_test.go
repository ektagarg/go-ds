package circularlinklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning

var cll = NewCircularLinkedList()

func TestInsertBeginning(t *testing.T) {
	node1 := NewNode(2)
	node2 := NewNode(34)
	node3 := NewNode(93)
	cll.InsertBeginning(node1)
	cll.InsertBeginning(node2)
	cll.InsertBeginning(node3)
	assert.Equal(t, 3, cll.Len, "InsertBeginning: Linked list length is incorrect!")

	// check if the first element is actually 93
	head := cll.Start
	assert.Equal(t, 93, head.Data, "Linkedlist InsertBeginning: didn't work!")
}

func TestInsertEnd(t *testing.T) {
	// insert in end
	node4 := NewNode(222)
	cll.InsertEnd(node4)
	assert.Equal(t, 4, cll.Len, "InsertEnd: Circular Linked list length is incorrect!")

	// check if the last element is actually 222
	head := cll.Start
	for {
		if head.Next == cll.Start {
			break
		}
		head = head.Next
	}
	assert.Equal(t, 222, head.Data, "Linkedlist InsertEnd: didn't work!")
}

func TestInsertLocation(t *testing.T) {
	//insert at a particular position
	node5 := NewNode(555)
	cll.InsertLocation(node5, 3)
	assert.Equal(t, 5, cll.Len, "InsertLocation: Circular Linked list length is incorrect!")

	count := 1
	head := cll.Start

	for {
		if count == 3 {
			assert.Equal(t, 555, head.Data, "Circular Linkedlist InsertLocation: didn't work!")
			break
		}
		head = head.Next
		count++
	}
}

func TestDeleteBeginning(t *testing.T) {
	// delete from beginnning
	deleted := cll.DeleteBeginning()
	assert.Equal(t, 4, cll.Len, "DeleteBeginning: Linked list length is incorrect!")
	assert.Equal(t, 93, deleted, "DeleteBeginning: Deleted Element is incorrect!")
}

func TestDeleteEnd(t *testing.T) {
	// delete from end
	deleted := cll.DeleteEnd()
	assert.Equal(t, 3, cll.Len, "DeleteEnd: Linked list length is incorrect!")
	assert.Equal(t, 222, deleted, "DeleteEnd: Deleted Element is incorrect!")
}

func TestDeleteFromPosition(t *testing.T) {
	// delete from any position
	deleted := cll.DeleteFromPosition(2)
	assert.Equal(t, 2, cll.Len, "DeleteFromPosition: Linked list length is incorrect!")
	assert.Equal(t, 555, deleted, "DeleteFromPosition: Deleted Element is incorrect!")
}
