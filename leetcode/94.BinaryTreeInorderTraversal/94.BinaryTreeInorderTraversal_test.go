package binarytreeinordertraversal

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)


func TestInorderTraversal(t *testing.T) {
	n := &TreeNode {
		Val: 10,	
		Left: &TreeNode{Val: 8},
		Right: &TreeNode{Val: 12},
	}
	assert.Equal(t, []int{8,10,12}, inorderTraversal(n))
	
}