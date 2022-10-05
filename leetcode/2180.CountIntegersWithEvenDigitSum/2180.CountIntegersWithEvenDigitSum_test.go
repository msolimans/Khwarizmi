package countintegerswithevendigitsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountEven(t *testing.T) {
    tests := []struct{
        In int 
        Exp int 
    }{
        {In: 30, Exp: 14},
        {In: 4, Exp: 2},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, countEven(test.In))
    }
}
