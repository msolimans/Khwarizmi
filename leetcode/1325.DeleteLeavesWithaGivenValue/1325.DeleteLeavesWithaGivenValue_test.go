package deleteleaveswithagivenvalue

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In *TreeNode
		In2 int 
        Exp *TreeNode
    }{
        {In: InitTreeNode([]*int{Int(1),Int(2),Int(3),Int(2),nil,Int(2),Int(4)}), In2: 2, Exp: InitTreeNodes(Int(1),nil,Int(3),nil,Int(4))},
        {In: InitTreeNode([]*int{Int(1),Int(3),Int(3),Int(3),Int(2)}), In2: 3, Exp: InitTreeNodes(Int(1),Int(3),nil,nil,Int(2))},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, removeLeafNodes(test.In, test.In2))
    }
}

