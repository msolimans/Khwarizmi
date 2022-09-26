package balancedbinarytree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct{
		In *TreeNode 
		Exp bool 
	} {
		{In: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}, Exp: true},
		{In: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 11}}}, Exp: false },
		{In: &TreeNode{Val: 3}, Exp: true },

	}

	for _,test:=range tests {
		assert.Equal(t, test.Exp, isBalanced(test.In))
	}
}