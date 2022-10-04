package findfirstandlastpositionofelementinsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
    tests := []struct{
        In []int 
		Target int 
        Exp []int 
    }{
        {In: []int{5,7,7,8,8,10}, Target: 8, Exp: []int{3,4}},
        {In: []int{5,7,7,8,8,10}, Target: 6, Exp: []int{-1,-1}},
        {In: []int{}, Target: 0, Exp: []int{-1,-1}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, searchRange(test.In, test.Target))
    }
}
