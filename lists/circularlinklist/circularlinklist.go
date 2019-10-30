package circularlinklist

import "fmt"

//Node ... newNode of circular linked list
type Node struct {
	Data int
	Next *Node
}

//CircularLinkedList ... A linked list with len, start
type CircularLinkedList struct {
	Len   int
	Start *Node
}

//NewNode ... Creating a new node
func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

//NewCircularLinkedList ... Creating an empty circular linked list
func NewCircularLinkedList() CircularLinkedList {
	return CircularLinkedList{
		Len:   0,
		Start: nil,
	}
}

//CheckIfEmpty ... check if empty
func (cll *CircularLinkedList) CheckIfEmpty() bool {
	if cll.Len == 0 {
		return true
	}
	return false
}

// CheckIfEmptyAndAdd ... check if doubly link list is empty
func (cll *CircularLinkedList) CheckIfEmptyAndAdd(newNode *Node) bool {
	// check if list is empty
	if cll.Len == 0 {
		// insert first node in doubly linked list
		cll.Start = newNode
		cll.Start.Next = newNode
		cll.Len++
		return true
	}
	return false
}

// InsertBeginning ... insert in the beginning
func (cll *CircularLinkedList) InsertBeginning(newNode *Node) {
	if !(cll.CheckIfEmptyAndAdd(newNode)) {
		head := cll.Start

		// insert on head
		newNode.Next = head

		// traverse this list until we find it's end
		for {
			if head.Next == cll.Start {
				break
			}
			head = head.Next
		}
		head.Next = newNode
		cll.Start = newNode
		cll.Len++
	}
}

//InsertEnd ... insert in the end
func (cll *CircularLinkedList) InsertEnd(newNode *Node) {
	if !(cll.CheckIfEmptyAndAdd(newNode)) {
		head := cll.Start
		for {
			if head.Next == cll.Start {
				break
			}
			head = head.Next
		}
		head.Next = newNode
		newNode.Next = cll.Start
		cll.Len++
	}
}

//InsertLocation ... Insert in the list on a specific location
func (cll *CircularLinkedList) InsertLocation(newNode *Node, pos int) {
	if !(cll.CheckIfEmptyAndAdd(newNode)) {
		head := cll.Start
		count := 1
		if pos == 1 {
			cll.InsertBeginning(newNode)
			return
		}
		for {
			if head.Next == nil && pos-1 > count {
				break
			}
			if count == pos-1 {
				newNode.Next = head.Next
				head.Next = newNode
				cll.Len++
				break
			}
			head = head.Next
			count++
		}
	}
}

//DeleteBeginning ... Delete from the start of the list
func (cll *CircularLinkedList) DeleteBeginning() int {
	//check if list if empty
	if !(cll.CheckIfEmpty()) {
		head := cll.Start
		deletedElem := head.Data
		if cll.Len == 1 {
			cll.Start = nil
			cll.Len--
			return deletedElem
		}
		prevStart := cll.Start
		cll.Start = head.Next
		// traverse till end and update last node's next to updated start
		for {
			if head.Next == prevStart {
				break
			}
			head = head.Next
		}
		head.Next = cll.Start
		cll.Len--
		return deletedElem
	}
	return -1
}

//DeleteEnd ... Delete the last element from circular list
func (cll *CircularLinkedList) DeleteEnd() int {
	if !(cll.CheckIfEmpty()) {
		head := cll.Start
		deletedEle := head.Data
		if cll.Len == 1 {
			// delete from beginning
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		//traverse till end
		for {
			if head.Next.Next == cll.Start {
				deletedEle = head.Next.Data
				break
			}
			head = head.Next
		}
		// update last element's next pointer
		head.Next = cll.Start
		cll.Len--
		return deletedEle
	}
	return -1
}

//DeleteFromPosition ... delete an element from circular list
func (cll *CircularLinkedList) DeleteFromPosition(pos int) int {
	if !(cll.CheckIfEmpty()) {
		head := cll.Start
		deletedEle := head.Data
		if cll.Len == 1 {
			// delete from beginning
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		if cll.Len == pos {
			// delete from end
			deletedEle = cll.DeleteEnd()
			return deletedEle
		}
		// delete from middle
		//traverse till you find position
		count := 1
		for {
			if count == pos-1 {
				deletedEle = head.Next.Data
				break
			}
			head = head.Next
		}
		head.Next = head.Next.Next
		cll.Len--
		return deletedEle
	}
	return -1
}

//PrintList ... Print Circular list
func (cll *CircularLinkedList) PrintList() {
	head := cll.Start
	for i := 0; i < cll.Len; i++ {
		fmt.Print(head.Data)
		fmt.Print("-->")
		head = head.Next
	}
	fmt.Println()
}
