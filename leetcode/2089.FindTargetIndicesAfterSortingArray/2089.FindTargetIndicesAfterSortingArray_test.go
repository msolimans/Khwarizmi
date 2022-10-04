package findtargetindicesaftersortingarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetIndices(t *testing.T) {
    tests := []struct{
        In []int 
		Target int 
        Exp []int 
    }{
        {In: []int{1,2,5,2,3}, Target: 2, Exp: []int{1,2} },
        {In: []int{1,2,5,2,3}, Target: 3, Exp: []int{3} },
        {In: []int{1,2,5,2,3}, Target: 5, Exp: []int{4} },
        {In: []int{1,2,5,2,3}, Target: 9, Exp: []int{} },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, targetIndices(test.In, test.Target))
    }
}
