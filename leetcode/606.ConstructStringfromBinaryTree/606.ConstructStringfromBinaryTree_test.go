package constructstringfrombinarytree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestTree2Str(t *testing.T) {
	tests := []struct{
		In *TreeNode
		Exp string
	}{
		{In: InitTreeNode([]*int{Int(1),Int(2),Int(3),Int(4)}), Exp: "1(2(4))(3)"},
		{In: InitTreeNode([]*int{Int(1),Int(2),Int(3),nil,Int(4)}), Exp: "1(2()(4))(3)"},
		{In: InitTreeNode([]*int{Int(1),nil,Int(2)}), Exp: "1()(2)"},
	}

	for _, test := range tests {
		assert.Equal(t, test.Exp, tree2str(test.In))
	}
}