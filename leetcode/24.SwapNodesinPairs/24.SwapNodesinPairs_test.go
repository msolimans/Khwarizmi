package swapnodesinpairs

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct{
        In *ListNode
        Exp *ListNode
    }{
        
        {In: Init([]int{1,2,3,4,5}) , Exp: Init([]int{2,1,4,3,5})},
        {In: Init([]int{1,2,3,4}) , Exp: Init([]int{2,1,4,3})},
        {In: Init([]int{1,2,3}) , Exp: Init([]int{2,1,3})},
        {In: Init([]int{1}) , Exp: Init([]int{1})},
        {In: Init([]int{}) , Exp: Init([]int{})},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, swapPairs(test.In))
    }
}