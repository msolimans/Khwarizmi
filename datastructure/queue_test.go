package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	var q Queue
	q.Enqueue("hello")
	q.Enqueue("world")
	assert.Equal(t, 2, q.Len())
	assert.False(t, q.IsEmpty())
}


func TestDequeue(t *testing.T) {
	var q Queue
	q.Enqueue("hello")
	q.Enqueue("world")
	
	assert.Equal(t, "hello", q.Dequeue())
	assert.Equal(t, "world", q.Dequeue())
}


