package poweroftwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool 
    }{
        {In: 5, Exp: false},
        {In: 4, Exp: true},
        {In: 8, Exp: true},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPowerOfTwo(test.In))
    }
}


func Test2(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool 
    }{
        {In: 5, Exp: false},
        {In: 4, Exp: true},
        {In: 8, Exp: true},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPowerOfTwo2(test.In))
    }
}

