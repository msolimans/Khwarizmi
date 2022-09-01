package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestLQueue_Ops(t *testing.T) {
	var l LLQueue
	l.Enqueue(10)
	l.Enqueue(20)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 10, l.Dequeue())
	assert.Equal(t, 20, l.Dequeue())
}