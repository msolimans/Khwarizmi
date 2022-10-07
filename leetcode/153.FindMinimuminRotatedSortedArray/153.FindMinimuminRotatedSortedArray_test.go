package findminimuminrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindM(t *testing.T) {
	tests := []struct{
        In []int  
        Exp int 
    }{
        {In:[] int{3,4,5,1,2}  , Exp: 1},
        {In:[] int{4,5,6,7,0,1,2}  , Exp: 0},
        {In:[] int{11,13,15,17}  , Exp: 11},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, findMin(test.In))
    }
}