package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOps(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	v, k := s.Peek()
	assert.Equal(t, 30, v)
	assert.True(t, k)

	vp,kp := s.Pop()
	assert.True(t, kp)
	assert.Equal(t, 30, vp)

	vp,kp = s.Pop()
	assert.True(t, kp)
	assert.Equal(t, 20, vp)

	vp,kp = s.Pop()
	assert.True(t, kp)
	assert.Equal(t, 10, vp)

	vp,kp = s.Pop()
	assert.False(t, kp)
	assert.Equal(t, nil, vp)

	vp,kp = s.Peek()
	assert.False(t, kp)
	assert.Equal(t, nil, vp)
}