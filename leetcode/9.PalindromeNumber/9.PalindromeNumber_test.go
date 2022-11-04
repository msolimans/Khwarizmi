package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp bool 
    }{
        {In: 121, Exp: true},
        {In: -121, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, isPalindrome(test.In))
    }
}