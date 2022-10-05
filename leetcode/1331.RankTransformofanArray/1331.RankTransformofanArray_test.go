package ranktransformofanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayRankTransform(t *testing.T) {
    tests := []struct{
        In []int 
        Exp []int 
    }{
        {In: []int{40,10,20,30}, Exp: []int{4,1,2,3} },
        {In: []int{100,100,100}, Exp: []int{1,1,1} },
        {In: []int{37,12,28,9,100,56,80,5,12}, Exp: []int{5,3,4,2,8,6,7,1,3} },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, arrayRankTransform(test.In))
    }
}
