package swappingnodesinalinkedlist

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestSwapNodes(t *testing.T) {
    tests := []struct{
        In *ListNode
        K int 
		Exp *ListNode
    }{
        {In: Init([]int{1,2,3,4,2,5}), K: 2, Exp: Init([]int{1,2,3,4,2,5}) },
        {In: Init([]int{1,2,3,4,2,5}), K: 1, Exp: Init([]int{5,2,3,4,2,1}) },
        {In: Init([]int{1, 2,3}), K: 1, Exp: Init([]int{3,2,1}) },
        {In: Init([]int{1,2,3}), K: 2, Exp: Init([]int{1,2,3}) },
        {In: Init([]int{1,2}), K: 1, Exp: Init([]int{2,1}) },
        {In: Init([]int{1,3,2,2}), K: 1, Exp: Init([]int{2,3,2,1}) },
        {In: Init([]int{1,3,2,2}), K: 2, Exp: Init([]int{1,2,3,2}) },
        {In: Init([]int{100,90}), K: 2, Exp: Init([]int{90,100}) },
		//0,1,2,3,4,2,5
		//f f f f f f f
		//s s s s s
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, swapNodes(test.In, test.K))
    }
}