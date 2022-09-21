package removeduplicatesfromsortedlist

import (
	"testing"

	"github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T){
	tests := []struct {
		In *common.ListNode 
		Expected *common.ListNode
	}{
		{
			In: common.Init([]int{1,1,2,2,3,3}),
			Expected: common.Init([]int{1,2,3}),
		},
		{
			In: common.Init([]int{1,1,1,2,2,3,3}),
			Expected: common.Init([]int{1,2,3}),
		},
		{
			In: common.Init([]int{1,1,1}),
			Expected: common.Init([]int{1}),
		},
		{
			In: common.Init([]int{}),
			Expected: common.Init([]int{}),
		},

	}

	for _, test := range tests {
		assert.True(t, deleteDuplicates(test.In).Equals(test.Expected))
	}
}