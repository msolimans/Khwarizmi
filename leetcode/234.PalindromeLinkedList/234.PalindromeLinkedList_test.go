package palindromelinkedlist

import (
	"testing"

	"github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
    tests := []struct{
        In *common.ListNode
        Exp bool 
    }{
        {In: common.Init([]int{1,2,2,1}) , Exp: true},
        {In: common.Init([]int{1,2,3,2,1}) , Exp: true},
        {In: common.Init([]int{1,2,1}) , Exp: true},
        {In: common.Init([]int{1,1}) , Exp: true},
        {In: common.Init([]int{1}) , Exp: true},
        {In: common.Init([]int{1,2}) , Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, isPalindrome(test.In))
    }
}