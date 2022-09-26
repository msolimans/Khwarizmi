package sametree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T){
	tree := &TreeNode{Val: 2, Left: &TreeNode{Val: 1},Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1},Right: &TreeNode{Val: 3}}
	assert.Equal(t, true, isSameTree(tree, tree2))

}