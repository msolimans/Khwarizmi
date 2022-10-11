package increasingtripletsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp bool 
    }{
        {In: []int{1,2,3,4,5}, Exp: true },
        {In: []int{5,4,3,2,1}, Exp: false },
        {In: []int{2,1,5,0,4,6}, Exp: true },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,increasingTriplet(test.In))
    }
}
