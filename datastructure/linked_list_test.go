package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T){
	list := Init()
	list.PushBack(10)
	list.PushBack(20)
	v, ok := list.PopFront()
	assert.True(t, ok)
	assert.Equal(t, 10, v)
	vl, ok := list.PopLast()
	assert.True(t, ok)
	assert.Equal(t, 20, vl)
	list.PushFront(50)
	list.PushFront(40)
	list.PushFront(30)
	vf, ok := list.PopLast()
	assert.True(t, ok)
	assert.Equal(t, 50, vf)
	
}