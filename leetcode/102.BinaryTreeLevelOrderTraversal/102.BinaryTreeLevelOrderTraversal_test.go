package binarytreelevelordertraversal

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {

	//add more test cases as needed
 
	tests := []struct{
		In *TreeNode 
		Exp [][]int
	} {
		{
			In:	&TreeNode{
					Val: 4,
					Left: &TreeNode {
						Val: 1,
					},
					Right: &TreeNode {
						Val: 6,
						Left: &TreeNode {Val: 5},
					},
				},
			Exp: [][]int{{4},{1,6}, {5}},
			
		},
		{
			In:	&TreeNode{
					Val: 4,
				},
			Exp: [][]int{{4}},
			
		},
		{
			In:	nil,
			Exp: [][]int{},
			
		},
	}
		
	for _, test:= range tests {
		assert.Equal(t, test.Exp, levelOrder(test.In))
	}

}