package pathsum

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct{
        In *TreeNode 
		TargetSum int 
        Exp bool 
    }{
        {In: InitTreeNode([]*int{}),TargetSum: 5, Exp: false },
        {In: InitTreeNode([]*int{Int(1)}),TargetSum: 5, Exp: false },
        {In: InitTreeNode([]*int{Int(1), Int(2), Int(3)}),TargetSum: 5, Exp: false },
        {In: InitTreeNode([]*int{Int(1), Int(2)}),TargetSum: 3, Exp: true },
        {In: InitTreeNode([]*int{Int(1), Int(2), nil, nil, Int(3)}),TargetSum: 6, Exp: true },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, hasPathSum(test.In, test.TargetSum))
    }
}