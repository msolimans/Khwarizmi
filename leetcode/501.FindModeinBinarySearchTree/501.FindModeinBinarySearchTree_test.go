package findmodeinbinarysearchtree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In *TreeNode
        Exp []int 
    }{
        {In: InitTreeNodes(Int(1),Int(1),Int(2),nil,nil,Int(2)) , Exp: []int{1,2} },
        {In: InitTreeNodes(Int(1),Int(1),Int(1),nil,nil,nil,Int(2)) , Exp: []int{1} },
        {In: InitTreeNodes(Int(0)) , Exp: []int{0} },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, findMode(test.In))
    }
}
