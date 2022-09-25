package reverselinkedlist

import (
	"testing"

	"github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	 
	exp := common.Init([]int{4,3,2,1})
	input := common.Init([]int{1,2,3,4})
	// assert.Equal(t, exp, reverseList(input))
	assert.Equal(t, exp, reverseList2(input))
}