package singlylinklist

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

func IsEmpty(list *LinkedList) bool {
	if list.Len == 0 {
		return true
	}
	return false
}

func InitializeEmpty(list *LinkedList, node *Node) {
	list.Start = node
	list.End = node
	list.Len++
	return
}

//InsertBeginning ... Insert at the start of the list
func (list *LinkedList) InsertBeginning(node *Node) {
	//check if list if empty
	if IsEmpty(list) {
		InitializeEmpty(list, node)
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
	if IsEmpty(list) {
		InitializeEmpty(list, node)
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
			list.Len++
			break
		}
		head = head.Next
	}
}

//InsertAtPosition ... Insert at a specific position
func (list *LinkedList) InsertAtPosition(node *Node, pos int) {
	//check if list if empty
	if IsEmpty(list) {
		InitializeEmpty(list, node)
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
			list.Len++
			break
		}
		currentLoc++
		head = head.Next
	}
}

//DeleteBeginning ... Delete from the start of the list
func (list *LinkedList) DeleteBeginning() int {
	//check if list if empty
	if list.Len == 0 {
		return 0
	} else if list.Len == 1 {
		head := list.Start
		list.Start = nil
		list.End = nil
		list.Len--
		return head.Data
	} else {
		head := list.Start
		list.Start = head.Next
		list.Len--
		return head.Data
	}
}

//DeleteEnd ... Delete from the end of the list
func (list *LinkedList) DeleteEnd() int {
	//check if list if empty
	if list.Len == 0 {
		return 0
	} else if list.Len == 1 {
		head := list.Start
		list.Start = nil
		list.End = nil
		list.Len--
		return head.Data
	} else {
		head := list.Start
		previous := list.Start
		//traverse till end
		for {
			if head.Next == nil {
				break
			}
			previous = head
			head = head.Next
		}
		previous.Next = nil
		list.Len--
		return previous.Data
	}
}

//DeleteFromPosition ... Delete from a specific position
func (list *LinkedList) DeleteFromPosition(pos int) int {
	//check if list if empty
	if list.Len == 0 {
		return 0
	} else if list.Len == 1 {
		head := list.Start
		list.Start = nil
		list.End = nil
		list.Len--
		return head.Data
	} else {
		head := list.Start
		previous := list.Start
		if pos == 1 {
			//call delete from Beginning
			return list.DeleteBeginning()
		}
		for i := 0; i < pos-1; i++ {
			previous = head
			head = previous.Next
		}
		previous.Next = head.Next
		head.Next = nil
		list.Len--
		return head.Data
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
