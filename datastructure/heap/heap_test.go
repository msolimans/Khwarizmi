package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapOps(t *testing.T) {
	h := Init()
	h.Push(10)
	h.Push(20)
	h.Push(5)
	assert.Equal(t, 5, h.Pop())
	h.Push(2)
	assert.Equal(t, 2, h.Top())
	h.Push(1)
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 1, h.Top())
	assert.Equal(t, 1, h.Pop())


}