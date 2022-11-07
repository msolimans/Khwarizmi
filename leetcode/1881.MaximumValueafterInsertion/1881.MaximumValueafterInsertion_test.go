package maximumvalueafterinsertion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
		X int 
        Exp string 
    }{
        {In: "99" , X: 9, Exp: "999"},
        {In: "-13" , X: 9, Exp: "-139"},
        {In: "-13" , X: 2, Exp: "-123"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,maxValue(test.In, test.X))
    }
}
