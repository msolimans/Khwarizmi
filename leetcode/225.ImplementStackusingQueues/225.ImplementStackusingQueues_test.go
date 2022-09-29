package implementstackusingqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOps(t *testing.T) {
	s := Constructor()
	s.Push(10)
	assert.Equal(t, 10, s.top)
	assert.Equal(t, 10, s.Pop())
	s.Push(20)
	s.Push(30)
	s.Push(40)
	assert.Equal(t, 40, s.Pop())
	assert.Equal(t, 30, s.Top())
	assert.Equal(t, 30, s.Pop())
	assert.Equal(t, 20, s.Pop())
	s.Push(20)
	assert.Equal(t, 20, s.Top())
	assert.Equal(t, 20, s.Pop())
	assert.Equal(t, -1, s.Top())
	assert.Equal(t, -1, s.Pop())
}
