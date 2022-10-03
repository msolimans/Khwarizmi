package quick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct{
        In []int 
        Exp []int 
    }{
        {In: []int{9,8,2,3,1,5,6,7} , Exp: []int{1,2,3,5,6,7,8,9}},
    }

    for _, test := range tests {
        sort(test.In) 
		assert.Equal(t, test.Exp, test.In)
    }
}