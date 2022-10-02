package minimumsumoffourdigitnumberaftersplittingdigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSum(t *testing.T) {
	tests := []struct{
        In int 
        Exp int 
    }{
        {In: 2932, Exp: 52},
        {In: 4009, Exp: 13},
    }

    for _, test := range tests {
		assert.Equal(t, test.Exp,  minimumSum(test.In))
    }
	
}