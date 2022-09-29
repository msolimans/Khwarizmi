package maximumtwinsumofalinkedlist

import (
	"testing"

	"github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	tests := []struct{
		In *common.ListNode
		Exp int 
	}{
		// {In: common.Init([]int{1,2,3,4}), Exp: 5},
		{In: common.Init([]int{1,2,2,3,9,4}), Exp: 11},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, pairSum(test.In))
	}
}

func TestPairSum2(t *testing.T) {
	tests := []struct{
		In *common.ListNode
		Exp int 
	}{
		{In: common.Init([]int{1,2,3,4}), Exp: 5},
		{In: common.Init([]int{1,2,2,3,9,4}), Exp: 11},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, pairSum2(test.In))
	}
}