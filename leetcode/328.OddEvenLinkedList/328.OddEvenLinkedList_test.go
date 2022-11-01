package OddEvenLinkedList

import (
	"testing"
	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In *ListNode 
        Exp *ListNode 
    }{
        {In: Init([]int{1,2,3,4,5}), Exp: Init([]int{1,3,5,2,4}) },
        {In: Init([]int{2,1,3,5,6,4,7}), Exp: Init([]int{2,3,6,7,1,5,4}) },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,oddEvenList(test.In))
    }
}

