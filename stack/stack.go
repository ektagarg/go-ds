package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	element []int
}

// Creates an empty integer stack
func NewStack() Stack {
	return Stack{element: make([]int, 0)}
}

//Check is array is empty
func (st *Stack) IsEmpty() bool {
	return len(st.element) <= 0
}

//returns top
func (st *Stack) Top() int {
	if st.IsEmpty() {
		return -1
	}
	return len(st.element) - 1
}

//returns size of array
func (st *Stack) Size() int {
	return len(st.element)
}

//push an element to stack
func (st *Stack) Push(element int) {
	st.element = append(st.element, element)
}

//pop an element
func (st *Stack) Pop() (removed int, err error) {
	if st.Top() != -1 {
		removed := st.element[st.Top()]
		st.element = st.element[:st.Size()-1]
		return removed, nil
	}
	return 0, errors.New("stack is empty")

}

// print elements in stack
func (st *Stack) Print() {
	for i := 0; i < st.Size(); i++ {
		fmt.Println("index, element:", i, st.element[i])
	}
}

func main() {
	st := NewStack()

	fmt.Println(st.Pop())

}
