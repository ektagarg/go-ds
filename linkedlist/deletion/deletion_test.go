package deletion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestDeletion ... test insert in the beginning
func TestDeletion(t *testing.T) {
	ll := NewLinkedList()
	n1 := NewNode(267567)
	n2 := NewNode(21)
	n3 := NewNode(32)
	n4 := NewNode(552)
	n5 := NewNode(223)

	ll.InsertBeginning(n1)
	ll.InsertBeginning(n2)
	ll.InsertBeginning(n3)
	ll.InsertBeginning(n4)
	ll.InsertBeginning(n5)

	ll.PrintList()

	ll.DeleteFromPosition(3)
	ll.PrintList()
	ll.DeleteBeginning()
	ll.PrintList()
	ll.DeleteEnd()
	ll.PrintList()

	assert.Equal(t, 3, ll.Len, "Linked list length is incorrect!")

}
