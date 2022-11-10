package validatebinarysearchtree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In *TreeNode
        Exp bool
    }{
        {In: InitTreeNodes(Int(2), Int(1), Int(3)) , Exp: true },
        {In: InitTreeNodes(Int(2), Int(2), Int(2)) , Exp: false },
        {In: InitTreeNodes(Int(1), nil, Int(1)) , Exp: false },
        {In: InitTreeNodes(Int(2) ) , Exp: true },
        {In: InitTreeNodes(Int(5),Int(1),Int(4),nil, Int(3), Int(6)) , Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isValidBST(test.In))
    }
}
