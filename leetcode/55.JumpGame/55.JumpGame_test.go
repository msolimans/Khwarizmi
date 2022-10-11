package jumpgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp bool 
    }{
        {In: []int{1,1,0,1}, Exp: false},
        {In: []int{2,3,1,1,4}, Exp: true},
        {In: []int{3,2,1,0,4}, Exp: false},
        {In: []int{1,2}, Exp: true},
        {In: []int{1,2,3}, Exp: true},
        {In: []int{3,2,1}, Exp: true},
        {In: []int{0,0}, Exp: false},
        {In: []int{0,1}, Exp: false},
        {In: []int{1,0}, Exp: true},
        {In: []int{1,0}, Exp: true},
        {In: []int{0}, Exp: true},

    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, canJump(test.In))
    }
}
