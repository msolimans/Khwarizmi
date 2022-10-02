package partitionlist

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
    tests := []struct{
        In *ListNode
		X int 
        Exp *ListNode
    }{
        {In: Init([]int{1,4,3,2,5,2}),X: 3, Exp: Init([]int{1,2,2,4,3,5})},
        {In: Init([]int{2,1}),X: 2, Exp: Init([]int{1,2})},
        {In: Init([]int{1}),X: 1, Exp: Init([]int{1})},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, partition(test.In, test.X))
    }
}