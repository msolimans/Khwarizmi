package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPermutation(t *testing.T) {
    tests := []struct{
        In []int 
        Exp [][]int 
    }{
        {In: []int{1} , Exp: [][]int{{1}}},
        {In: []int{1,2} , Exp: [][]int{{1,2},{2,1}}},
        {In: []int{1,2,3} , Exp: [][]int{{1,2,3},{1,3,2}, {2,1,3},{2,3,1},{3,2,1},{3,1,2}}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, permute(test.In))
    }
}
