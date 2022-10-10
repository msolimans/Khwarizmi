package maximumdifferencebetweenincreasingelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp int 
    }{
        {In: []int{9,4,3,2} , Exp: -1 },
        {In: []int{7,1,5,4} , Exp: 4 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, maximumDifference(test.In))
    }
}