package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLLStackOperations(t *testing.T) {
	stack := InitLLStack() 
	stack.Push(10)
	stack.Push(20)


	v20, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t,20, v20)

	v10, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t,10, v10)
	
	assert.True(t, stack.IsEmpty())

}