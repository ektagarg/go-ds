package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = NewTree()

func TestInsert(t *testing.T) {
	tree.Insert(34)
	tree.Insert(67)
	tree.Insert(23)
	tree.Insert(43)
	tree.Insert(11)
	tree.Insert(22)
	tree.Insert(55)
}

//TestPreOrder ... Test pre order traversal in binary search tree
func TestPreOrder(t *testing.T) {
	traversed := []int{}

	tree.PreOrder(tree.Root, &traversed)

	expected := []int{34, 23, 11, 22, 67, 43, 55}

	assert.Equal(t, 7, len(traversed), "Total number of nodes are not correct!")
	assert.Equal(t, expected, traversed, "pre order traversal didn't work")
}

//TestInOrder ... Test in order traversal in binary search tree
func TestInOrder(t *testing.T) {
	traversed := []int{}

	tree.InOrder(tree.Root, &traversed)

	expected := []int{11, 22, 23, 34, 43, 55, 67}

	assert.Equal(t, 7, len(traversed), "Total number of nodes are not correct!")
	assert.Equal(t, expected, traversed, "in order traversal didn't work")
}

//TestPostOrder ... Test post order traversal in binary search tree
func TestPostOrder(t *testing.T) {
	traversed := []int{}

	tree.PostOrder(tree.Root, &traversed)

	expected := []int{22, 11, 23, 55, 43, 67, 34}

	assert.Equal(t, 7, len(traversed), "Total number of nodes are not correct!")
	assert.Equal(t, expected, traversed, "post order traversal didn't work")
}
