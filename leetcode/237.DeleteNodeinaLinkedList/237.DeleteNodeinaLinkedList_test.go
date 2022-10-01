package deletenodeinalinkedlist

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Head *ListNode 
    In int 
    Exp *ListNode 
}

func (t *TestCase) getTestNode() *ListNode {
	c := t.Head
	for c!= nil {
		if c.Val == t.In {
			return c 
		}
		c = c.Next
	}
	return nil 

}
func TestDeleteNode(t *testing.T) {
    tests := []*TestCase{}

	tests = append(tests,  &TestCase{Head: Init([]int{1,2,3,4,5}), In: 2, Exp: Init([]int{1,3,4,5})})
	tests = append(tests,  &TestCase{Head: Init([]int{1,2,3,4,5}), In: 3, Exp: Init([]int{1,2,4,5})})

    for _, test := range tests {
        deleteNode(test.getTestNode()) 
		assert.Equal(t, test.Exp, test.Head)
    }
}