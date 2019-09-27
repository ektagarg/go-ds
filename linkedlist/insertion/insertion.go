package insertion

import "fmt"

//Node ... A node that conatins data and link to point to it's next node
type Node struct {
	Data int
	Next *Node
}

//LinkedList ... A linked list with len, start and end
type LinkedList struct {
	Len   int
	Start *Node
	End   *Node
}

//NewNode ... Creating a new node
func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

//NewLinkedList ... Creating an empty linked list
func NewLinkedList() LinkedList {
	return LinkedList{
		Len:   0,
		Start: nil,
		End:   nil,
	}
}

//InsertBeginning ... Insert at the start of the list
func (list *LinkedList) InsertBeginning(node *Node) {
	//check if list if empty
	if list.Len == 0 {
		list.Start = node
		list.End = node
		list.Len++
		return
	}

	head := list.Start
	list.Start = node
	node.Next = head
	list.Len++
}

//InsertEnd ... Insert at the end of the list
func (list *LinkedList) InsertEnd(node *Node) {
	//check if list if empty
	if list.Len == 0 {
		list.Start = node
		list.End = node
		list.Len++
		return
	}

	head := list.Start
	//traverse till end
	for {
		if head.Next == nil {
			// here we need to insert
			head.Next = node
			node.Next = nil
			list.End = node
			break
		}
		head = head.Next
	}
}

//InsertAtPosition ... Insert at a specific position
func (list *LinkedList) InsertAtPosition(node *Node, pos int) {
	//check if list if empty
	if list.Len == 0 {
		list.Start = node
		list.End = node
		list.Len++
		return
	}

	head := list.Start
	currentLoc := 1

	if pos == 1 {
		//call Insert at Beginning
		list.InsertBeginning(node)
		return
	}
	//traverse till that specific position
	for {
		if head.Next == nil && pos-1 > currentLoc {
			break
		}
		if currentLoc == pos-1 {
			// here we need to insert
			node.Next = head.Next
			head.Next = node
			break
		}
		currentLoc++
		head = head.Next
	}
}

//TraverseAndPrint ... Traverse and print linked list
func (list *LinkedList) TraverseAndPrint() {
	head := list.Start
	for {
		fmt.Print(head.Data)
		if head.Next == nil {
			break
		}
		head = head.Next
		fmt.Print("->")
	}
}
