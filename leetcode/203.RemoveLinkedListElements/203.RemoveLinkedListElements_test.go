package removelinkedlistelements

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		In *ListNode
		In2 int 
		Exp *ListNode
	} {
		{In: Init([]int{5,1,2,3,4,5,5,5,6,5}), In2: 5, Exp: Init([]int{1,2,3,4,6})},
		{In: Init([]int{5,5,5,5}), In2: 5, Exp: nil},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, removeElements(test.In, test.In2))
	}


}