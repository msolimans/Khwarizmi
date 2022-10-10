package twofurthesthouseswithdifferentcolors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp int 
    }{
        {In: []int{1,1,1,6,1,1,1} , Exp: 3},
        {In: []int{1,8,3,8,3} , Exp: 4},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,maxDistance(test.In))
    }
}