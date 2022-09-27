package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	stack := Constructor() 
	stack.Push(2147483646)
	stack.Push(2147483646)
	stack.Push(2147483647)
	assert.Equal(t,2147483647, stack.Top())
	stack.Pop()
	assert.Equal(t,2147483646, stack.GetMin())
	stack.Pop()
	assert.Equal(t,2147483646, stack.GetMin())
	stack.Pop()
	stack.Push(2147483647)
	assert.Equal(t,2147483647, stack.Top())
	assert.Equal(t,2147483647, stack.GetMin())

}

