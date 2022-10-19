package countingbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp []int 
    }{
        {In: 5, Exp: []int{0, 1, 1, 2, 1, 2}},
        {In: 10, Exp: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,countBits(test.In))
    }
}


func Test2(t *testing.T) {
    tests := []struct{
        In int 
        Exp []int 
    }{
        {In: 5, Exp: []int{0, 1, 1, 2, 1, 2}},
        {In: 10, Exp: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,countBits2(test.In))
    }
}
