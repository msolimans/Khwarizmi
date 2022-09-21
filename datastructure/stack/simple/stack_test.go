package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T){
	var s Stack 
	s.Push("hello")
	s.Push("world")

	assert.Equal(t, 2, s.Len())	
}

func TestPopt(t *testing.T) {
	var s Stack 
	s.Push("hello")
	s.Push("world")

	for !s.IsEmpty() { 
		assert.Equal(t, "world", s.Pop())
		assert.Equal(t, "hello", s.Pop())
	}

}