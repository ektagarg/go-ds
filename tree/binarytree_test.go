package tree

import (
	"testing"
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
	tree.InOrder(tree.Root)
}
