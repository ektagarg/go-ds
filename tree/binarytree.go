package tree

import (
	"errors"
	"fmt"
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
	} else {
		if node.Left == nil {
			node.Left = NewNode(data)
			return nil
		}
	}
	return node.Left.InsertNode(data)
}

func (t *Tree) InOrder(node *Node) error {
	if node == nil {
		return nil
	}
	traversed := []int{}
	// process left subtree

	if node.Left != nil {
		fmt.Println("left......")
		fmt.Println(node.Left.Data)
		t.InOrder(node.Left)
	} else {
		traversed = append(traversed, node.Left.Data)
		fmt.Println("isssue here......")
	}

	//process root
	traversed = append(traversed, node.Data)

	// process right subtree
	if node.Right != nil {
		fmt.Println("right......")
		fmt.Println(node.Right.Data)
		t.InOrder(node.Right)
	} else {
		traversed = append(traversed, node.Right.Data)
	}

	fmt.Println("traversed......")
	fmt.Println(traversed)
	return nil
}
