package twosumivinputisabst

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestFindTarget(t *testing.T) {
	tests := []struct{
        In *TreeNode
		K int 
        Exp bool
    }{
        {In: InitTreeNode([]*int{Int(5), Int(3),Int(6),Int(2),Int(4),nil,Int(7)}), K: 9, Exp: true},
        {In: InitTreeNode([]*int{Int(5), Int(3),Int(6),Int(2),Int(4),nil,Int(7)}), K: 28, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, findTarget(test.In,test.K))
    }
}



func TestFindTarget2(t *testing.T) {
	tests := []struct{
        In *TreeNode
		K int 
        Exp bool
    }{
        {In: InitTreeNode([]*int{Int(5), Int(3),Int(6),Int(2),Int(4),nil,Int(7)}), K: 9, Exp: true},
        {In: InitTreeNode([]*int{Int(5), Int(3),Int(6),Int(2),Int(4),nil,Int(7)}), K: 28, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, findTarget2(test.In,test.K))
    }
}