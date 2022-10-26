package medianoftwosortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test(t *testing.T) {
    tests := []struct{
        Arr1 []int 
        Arr2 []int 
		Exp float64
    }{
        {Arr1: []int{1,3}, Arr2: []int{2}, Exp: 2.0},
    }

    for _, test := range tests {
        assert.Equal(t, 2.0 ,findMedianSortedArrays(test.Arr1, test.Arr2))
    }
}
