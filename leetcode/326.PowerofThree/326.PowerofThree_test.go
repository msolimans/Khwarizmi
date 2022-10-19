package powerofthree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool 
    }{
        {In: 0, Exp: false},
        {In: -1, Exp: false},
        {In: 3, Exp: true},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPowerOfThree(test.In))
    }
}

func Test2(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool 
    }{
        {In: 0, Exp: false},
        {In: -1, Exp: false},
        {In: 3, Exp: true},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPowerOfThree2(test.In))
    }
}

