package lowestcommonancestorofabinarysearchtree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

type TestCase struct{
	In *TreeNode
	In2 *TreeNode
	In3 *TreeNode 
	Exp *TreeNode 
}

func Case1() *TestCase {
	tree := InitTreeNode([]*int{Int(6),Int(2),Int(8),Int(0),Int(4),Int(7),Int(9),nil,nil,Int(3),Int(5)}) 
	return &TestCase{
		In: tree,
		In2: tree.Left,
		In3: tree.Right,
		Exp: tree,
	}
}

func Case2() *TestCase {
	tree := InitTreeNode([]*int{Int(6),Int(2),Int(8),Int(0),Int(4),Int(7),Int(9),nil,nil,Int(3),Int(5)}) 
	return &TestCase{
		In: tree,
		In2: tree.Left.Right,
		In3: tree.Left,
		Exp: tree.Left,
	}
}

func TestLowestCommonAncestor(t *testing.T) {
    tests := []*TestCase{Case1(), Case2()}
       
    for _, test := range tests {
        assert.Equal(t, test.Exp, lowestCommonAncestor(test.In, test.In2, test.In3))
    }
}