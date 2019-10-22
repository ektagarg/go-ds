package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestInsertion ... test insert in the beginning

var stack = NewStack()

func TestPush(t *testing.T) {
	stack.Push(923)
	stack.Push(213)
	stack.Push(834)
	stack.Push(443)
	assert.Equal(t, 4, stack.Size(), "Stack size incorrect!")
}

func TestPop(t *testing.T) {
	x, _ := stack.Pop()
	assert.Equal(t, 3, stack.Size(), "Stack size incorrect!")
	assert.Equal(t, 443, x, "Top element didn't get pop!")
}
