package makearrayzerobysubtractingequalamounts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOperations(t *testing.T) {
	tests := []struct{
        In []int 
        Exp int
    }{
        {In: []int{1,5,0,3,5}, Exp: 3},
        {In: []int{0}, Exp: 0},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, minimumOperations(test.In))
    }
}