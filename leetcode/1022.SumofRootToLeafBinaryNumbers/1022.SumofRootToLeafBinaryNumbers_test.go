package sumofroottoleafbinarynumbers

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
        {In: InitTreeNodes(Int(1),Int(0),Int(1),Int(0),Int(1),Int(0),Int(1)) , Exp: 22 },
        {In: InitTreeNodes(Int(0)) , Exp: 0 },

    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,sumRootToLeaf(test.In))
    }
}
