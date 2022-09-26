package deletethemiddlenodeofalinkedlist

import (
	"testing"

	"github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddle(t *testing.T){
	assert.Equal(t, common.Init([]int{1,2,4}), deleteMiddle(common.Init([]int{1,2,3,4})))
}