package doublylinklist

import "fmt"

//Node ... newNode of doubly linked list
type Node struct {
	Data int
	Prev *Node // points to previous node in link list
	Next *Node // points to next node in link list
}

//Doublylinklist ...
type Doublylinklist struct {
	Length int
	Start  *Node
	End    *Node
}

//NewNode ... it return a new node
func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}

//NewDoublylinklist ... returns a doubly linked list
func NewDoublylinklist() Doublylinklist {
	return Doublylinklist{
		Length: 0,
		Start:  nil,
		End:    nil,
	}
}

//CheckIfEmpty ... check if empty
func (dll *Doublylinklist) CheckIfEmpty() bool {
	if dll.Length == 0 {
		return true
	}
	return false
}

// CheckIfEmptyAndAdd ... check if doubly link list is empty
func (dll *Doublylinklist) CheckIfEmptyAndAdd(newNode *Node) bool {
	// check if list is empty
	if dll.Length == 0 {
		// insert first node in doubly linked list
		dll.Start = newNode
		dll.End = newNode
		dll.Length++
		return true
	}
	return false
}

//InsertBeginning ... insert in the beginning of doubly linked list
func (dll *Doublylinklist) InsertBeginning(newNode *Node) {
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.Start
		// update newnode links - prev and next
		newNode.Next = head
		newNode.Prev = nil

		// update head node
		head.Prev = newNode

		// update dll start and length
		dll.Start = newNode
		dll.Length++
		return
	}
	return
}

//InsertEnd ... inserts a node in the end of doubly linked list
func (dll *Doublylinklist) InsertEnd(newNode *Node) {
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.Start
		for i := 0; i < dll.Length; i++ {
			if head.Next == nil {
				// update newnode links - prev and next
				newNode.Prev = head
				newNode.Next = nil

				//update head node
				head.Next = newNode

				// update dll end and Length
				dll.End = newNode
				dll.Length++
				break
			}
			head = head.Next
		}
	}
	return
}

// InsertMiddle ... insert between two nodes in doubly linkedlist
func (dll *Doublylinklist) InsertMiddle(newNode *Node, loc int) {
	if !(dll.CheckIfEmptyAndAdd(newNode)) {
		head := dll.Start
		for i := 1; i < dll.Length; i++ {
			if i == loc {
				// update newnode links - prev and next
				newNode.Prev = head.Prev
				newNode.Next = head

				// update head node
				head.Prev.Next = newNode
				head.Prev = newNode

				//keep traversing till we find the location
				dll.Length++
				return
			}
			head = head.Next
		}
	}
	return
}

//deleteFirst ... Delete last element
func (dll *Doublylinklist) DeleteFirst() int {
	if !(dll.CheckIfEmpty()) {
		head := dll.Start
		if head.Prev == nil {
			deletedNode := head.Data

			// update doubly linked list
			dll.Start = head.Next
			dll.Start.Prev = nil
			dll.Length--

			return deletedNode
		}
	}
	return -1
}

//DeleteLast ... deletes last element from doubly linked list
func (dll *Doublylinklist) DeleteLast() int {
	if !(dll.CheckIfEmpty()) {
		// delete from last
		head := dll.Start
		for {
			if head.Next == nil {
				break
			}
			head = head.Next
		}

		// update doubly linked list
		dll.End = head.Prev
		dll.End.Next = nil
		dll.Length--
		return head.Data
	}
	return -1
}

//DeleteMiddle .. delete from any location
func (dll *Doublylinklist) DeleteMiddle(pos int) int {
	if !(dll.CheckIfEmpty()) {
		//delete from any position
		head := dll.Start
		for i := 1; i <= pos; i++ {
			if head.Next == nil && pos > i {
				//list is lesser than given position
				return -1
			} else if i == pos {
				// delete from this location
				head.Prev.Next = head.Next
				head.Next.Prev = head.Prev
				dll.Length--
				return head.Data
			}
			head = head.Next
		}
	}
	return -1
}

//Print ... Print the doubly linked list
func (dll *Doublylinklist) Print() {
	head := dll.Start
	for i := 0; i < dll.Length; i++ {
		fmt.Println("-----newNode no------", i+1)
		fmt.Println("data", head.Data)
		fmt.Println("prev", head.Prev)
		fmt.Println("next", head.Next)
		fmt.Println("------------------")
		head = head.Next
	}
}
