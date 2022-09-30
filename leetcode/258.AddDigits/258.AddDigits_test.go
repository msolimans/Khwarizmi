package adddigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    tests := []struct{
        In int 
        Exp int 
    }{
        {In: 10 , Exp: 1 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, addDigits(test.In))
    }
}