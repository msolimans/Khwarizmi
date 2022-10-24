package poweroffour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool
    }{
        {In: 8, Exp: false},
        {In: 4, Exp: true},
        {In: 1, Exp: true},
        {In: 2, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPowerOfFour(test.In))
    }
}

