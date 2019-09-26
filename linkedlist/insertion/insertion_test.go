package insertion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning
func TestInsertion(t *testing.T) {
	linkedlist := NewLinkedList()
	node1 := NewNode(2)
	node2 := NewNode(34)
	node3 := NewNode(93)
	linkedlist.InsertBeginning(node1)
	linkedlist.InsertBeginning(node2)
	linkedlist.InsertBeginning(node3)
	assert.Equal(t, 3, linkedlist.Len, "Linked list length is incorrect!")

	// check if the first element is actually 93
	head := linkedlist.Start
	assert.Equal(t, 93, head.Data, "Linked list length insert at beginning didn't work!")

	// // insert in end
	node4 := NewNode(222)
	linkedlist.InsertEnd(node4)

	// check if the last element is actually 222
	end := linkedlist.End
	assert.Equal(t, 222, end.Data, "Linked list length insert at end didn't work!")

	//insert at a particular position
	node5 := NewNode(555)
	linkedlist.InsertAtPosition(node5, 3)
	count := 1
	for {
		fmt.Println(head.Data)
		if count == 3 {
			assert.Equal(t, 555, end.Data, "Linked list length insert at end didn't work!")
			break
		}
		head = head.Next
	}

}
