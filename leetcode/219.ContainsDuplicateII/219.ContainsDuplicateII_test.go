package containsduplicateii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
    tests := []struct{
        In []int
		In2 int
        Exp bool
    }{
        {In: []int{1,2,3,1},In2: 3, Exp: true},
        {In: []int{1,3,9,4,2,1,2,1},In2: 2, Exp: true},
        {In: []int{1,0,1,1},In2: 1, Exp: true},
        {In: []int{2,3,1},In2: 1, Exp: false},
        {In: []int{1,3,2,1,2,1},In2: 1, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, containsNearbyDuplicate(test.In, test.In2))
    }
}