package canplaceflowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
		N int 
        Exp bool 
    }{
        {In: []int{1} , N: 0, Exp: true},
        {In: []int{0} , N: 2, Exp: false},
        {In: []int{0} , N: 1, Exp: true},
        {In: []int{0,0} , N: 1, Exp: true},
        {In: []int{0,0} , N: 2, Exp: false},
        {In: []int{0,1} , N: 1, Exp: false},
        {In: []int{1,0,0,0,1} , N: 1, Exp: true},
        {In: []int{1,0,0,0,1} , N: 2, Exp: false},
        {In: []int{0,0,1,0,0,1} , N: 2, Exp: false},

    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,canPlaceFlowers(test.In,test.N))
    }
}

