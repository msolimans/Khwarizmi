package minimumabsolutedifferenceinbst

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In *TreeNode
        Exp int 
    }{
        {In: InitTreeNodes(Int(4),Int(2),Int(6),Int(1),Int(3)), Exp: 1},
        {In: InitTreeNodes(Int(4),Int(0),Int(48), nil, nil,Int(12),Int(49)), Exp: 1},
        {In: InitTreeNodes(Int(4),Int(2),Int(10), nil, nil,Int(6),nil), Exp: 2},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,getMinimumDifference(test.In))
    }
}
