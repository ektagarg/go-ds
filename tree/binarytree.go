package tree

import (
	"errors"
)

//Node ... create a new node for tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//Tree ...  binary tree has a root
type Tree struct {
	Root *Node
}

//NewNode ... Create a new node
func NewNode(data int) *Node {
	return &Node{data, nil, nil}
}

//NewTree ... create a new tree
func NewTree() *Tree {
	return &Tree{nil}
}

//Insert ... insert root in tree
func (t *Tree) Insert(data int) error {
	if t.Root == nil {
		t.Root = NewNode(data)
		return nil
	}
	return t.Root.InsertNode(data)
}

//InsertNode ... Insert a node in tree - left or right
func (node *Node) InsertNode(data int) error {
	if node == nil {
		return errors.New("Can't insert data to root")
	}

	if data > node.Data {
		if node.Right == nil {
			node.Right = NewNode(data)
			return nil
		}
		return node.Right.InsertNode(data)
	}
	if node.Left == nil {
		node.Left = NewNode(data)
		return nil
	}
	return node.Left.InsertNode(data)
}

//PreOrder ... root, left, right
func (t *Tree) PreOrder(node *Node, traversed *[]int) {
	if node == nil {
		return
	}
	// process root
	*traversed = append(*traversed, node.Data)

	// process left subtree
	t.PreOrder(node.Left, traversed)

	// process right subtree
	t.PreOrder(node.Right, traversed)
}

//InOrder ... left, root, right
func (t *Tree) InOrder(node *Node, traversed *[]int) {
	if node == nil {
		return
	}
	// process left subtree
	t.InOrder(node.Left, traversed)

	// process root
	*traversed = append(*traversed, node.Data)

	// process right subtree
	t.InOrder(node.Right, traversed)
}

//PostOrder ... left, right, root
func (t *Tree) PostOrder(node *Node, traversed *[]int) {
	if node == nil {
		return
	}
	// process left subtree
	t.PostOrder(node.Left, traversed)

	// process right subtree
	t.PostOrder(node.Right, traversed)

	// process root
	*traversed = append(*traversed, node.Data)
}
