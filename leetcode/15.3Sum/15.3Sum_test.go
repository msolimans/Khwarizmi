package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp [][]int 
    }{
        {In: []int{-1,0,1,2,-1,-4}, Exp: [][]int{{-1,-1,2}, {-1,0,1}}},
        {In: []int{0,1,1}, Exp: [][]int{}},
        {In: []int{0,0,0}, Exp: [][]int{{0,0,0}}},
        {In: []int{0,0,0,0,0}, Exp: [][]int{{0,0,0}}},
        {In: []int{-2,0,1,1,2}, Exp: [][]int{{-2, 0, 2},{-2, 1, 1}}},
        {In: []int{1,-1,-1,0}, Exp: [][]int{{-1, 0, 1}}},
        {In: []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}, Exp: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, threeSum(test.In))
    }
}
