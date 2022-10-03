package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
    tests := []struct{
        In []int 
        Exp int 
    }{
        {In: []int{3,2,3}, Exp: 3},
        {In: []int{2,2,1,1,1,2,2}, Exp: 2},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, majorityElement(test.In))
    }
} 