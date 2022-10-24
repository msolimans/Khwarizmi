package median

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        Arr1 []int 
        Arr2 []int 
      
    }{
        {Arr1: []int{10, 20, 30, 40, 51, 61, 71}, Arr2: []int{15, 25, 31, 86, 600, 700, 900}},
        {Arr1: []int{15, 25, 31, 86, 600, 700, 900}, Arr2: []int{10, 20, 30, 40, 51, 61, 71}},
        {Arr1: []int{15, 25, 31, 86, 600}, Arr2: []int{10, 20, 30, 40, 51}},
        {Arr1: []int{15, 25, 31, 86, 600}, Arr2: []int{10, 20, 30, 40, 51}}, //min => min 
		{Arr1: []int{15, 25, 31, 86}, Arr2: []int{10, 20, 30, 40}},//greater => greater 
        {Arr1: []int{15, 25, 31, 86, 600, 700}, Arr2: []int{10, 20, 30, 40, 51, 61}},
        {Arr1: []int{15, 25, 31, 86, 600}, Arr2: []int{10, 20, 30, 40, 51}},
        {Arr1: []int{15, 25, 31, 86}, Arr2: []int{10, 20, 30, 40}},
        {Arr1: []int{15, 25}, Arr2: []int{10, 20}},
        {Arr1: []int{15}, Arr2: []int{10}},
        {Arr1: []int{2,4,6,10}, Arr2: []int{8,9,10,12}},
        {Arr1: []int{2,4,6}, Arr2: []int{8,9,10}},
        {Arr1: []int{2,4}, Arr2: []int{8,9}},
        {Arr1: []int{2}, Arr2: []int{8}},
		//10,15,20,25,30,31,40,51,61,86,600,700 => 31,40
    }

    for _, test := range tests {
        assert.Equal(t, calcMedian2(test.Arr1, test.Arr2),calcMedian(test.Arr1, test.Arr2))
    }
}
