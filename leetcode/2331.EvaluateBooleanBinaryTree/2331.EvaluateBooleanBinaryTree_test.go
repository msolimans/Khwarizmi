package evaluatebooleanbinarytree

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateTree(t *testing.T){

	//2 means OR 
	//3 means AND 
	//0 means false 
	//1 means true 

	tests := []struct{
		In *TreeNode 
		Exp bool 
	}{
		{In: InitTreeNode([]*int{Int(1), nil}), Exp: true},
		{In: InitTreeNode([]*int{Int(2), Int(0), Int(1)}), Exp: true},
		{In: InitTreeNode([]*int{Int(2), Int(3), Int(2), Int(1), Int(0), Int(1), Int(0)}), Exp: true},
		{In: InitTreeNode([]*int{Int(3), Int(3), Int(2), Int(1), Int(0), Int(1), Int(0)}), Exp: false},
	}

	for _,test :=range tests {
		assert.Equal(t, test.Exp, evaluateTree(test.In))
	}
}