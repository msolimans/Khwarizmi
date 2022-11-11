package findclosestnumbertozero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp int 
    }{
        {In: []int{-4,-2,1,4,8}, Exp: 1 },
        {In: []int{2,-1,1}, Exp: 1 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,findClosestNumber(test.In))
    }
}
