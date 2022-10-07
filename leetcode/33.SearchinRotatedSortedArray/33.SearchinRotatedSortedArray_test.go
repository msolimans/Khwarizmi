package searchinrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
    tests := []struct{
        Nums []int
		Target int
        Exp int 
    }{
        {Nums: []int{4,5,6,7,0,1,2}, Target:0, Exp: 4 },
        {Nums: []int{4,5,6,7,0,1,2}, Target:3, Exp: -1 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, search(test.Nums, test.Target))
    }
}
