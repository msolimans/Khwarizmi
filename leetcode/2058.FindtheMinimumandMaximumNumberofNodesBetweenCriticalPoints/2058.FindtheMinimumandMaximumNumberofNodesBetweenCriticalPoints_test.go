package findtheminimumandmaximumnumberofnodesbetweencriticalpoints

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
    tests := []struct{
        In *ListNode
        Exp []int
    }{
        {In:Init([]int{5,3,1,2,5,1,2}), Exp:  []int{1,3} },
        {In:Init([]int{1,3,2,2,3,2,2,2,7}), Exp:  []int{3,3} },
        {In:Init([]int{3,1}), Exp:  []int{-1,-1} },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, nodesBetweenCriticalPoints(test.In))
    }
}