package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitTreeNode(t *testing.T) {
	
	// assert.NotNil(t, InitTreeNode([]*int{Int(1), Int(2), nil,}))
	assert.Nil(t, InitTreeNode([]*int{}))
	assert.NotNil(t, InitTreeNode([]*int{Int(1), Int(2), Int(3), nil,Int(4)}))
	tree1 := InitTreeNode([]*int{Int(1), Int(2), Int(3), nil,})
	assert.Equal(t, "1-2-3-nil-nil-nil-nil-", tree1.String())
}

func TestInt(t *testing.T) {
	i := Int(10)
	assert.Equal(t, 10, *i)
}