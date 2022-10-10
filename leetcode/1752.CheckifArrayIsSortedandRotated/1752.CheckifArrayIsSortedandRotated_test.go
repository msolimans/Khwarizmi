package checkifarrayissortedandrotated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp bool
    }{
        {In: []int{3,4,5,1,2}, Exp:true},
        {In: []int{2,1,3,4}, Exp:false},
        {In: []int{1,2,3}, Exp:true},
        {In: []int{5,6,4,2}, Exp:false},
        {In: []int{5,6,4,6}, Exp:false},
        {In: []int{7,9,1,1,1,3}, Exp:true},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,check(test.In))
    }
}