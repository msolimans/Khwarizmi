package narytreepreordertraversal

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestPreorder(t *testing.T) {
	tests := []struct{
		In *Node  
		Exp []int 
	}{
		{
			In: &Node{
				Val: 1,
				Children:[]*Node{{Val: 2}, {Val: 3}},
			},
			Exp: []int{1,2,3},
		},
	}
	for _,test := range tests{
		assert.Equal(t, test.Exp, preorder(test.In))
	}

	
}