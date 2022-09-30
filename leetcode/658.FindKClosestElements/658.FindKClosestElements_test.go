package findkclosestelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
 


func TestFindClosestElements(t *testing.T) {
	tests := []struct{
        Arr []int 
		K int 
		X int
        Exp []int 
    }{
        {Arr: []int{1,2,3,4,5},K: 4, X: 3, Exp: []int{1,2,3,4}},
        {Arr: []int{1,2,3,4,5},K: 1, X: 3, Exp: []int{3}},
        {Arr: []int{1,2,4,5},K: 1, X: 3, Exp: []int{2}},
        {Arr: []int{1,2,3,4,5},K: 4, X: -1, Exp: []int{1,2,3,4}},
        {Arr: []int{1,2,3,4,5},K: 1, X: -1, Exp: []int{1}},
        {Arr: []int{1,5},K: 1, X: -1, Exp: []int{1}},
        {Arr: []int{1,5},K: 1, X: 6, Exp: []int{5}},
        {Arr: []int{1,10,15,25,35,45,50,59},K: 1, X: 30, Exp: []int{25}},
	}
      
  
    for _, test := range tests {
        assert.Equal(t, test.Exp, findClosestElements(test.Arr, test.K, test.X))
    }
}


func TestFindClosestIndex(t *testing.T) {
	tests := []struct{
        Arr []int 
		K int 
		L int 
		R int 
        Exp int 
    }{
        {Arr: []int{1,2,3,4,5},K: 3, L:0,R:4, Exp: 2},
        {Arr: []int{1,2,5,5},K: 3, L:0,R:3, Exp: 1},
	}
      
  
    for _, test := range tests {
        assert.Equal(t, test.Exp, findClosestIndex(test.Arr, test.K, test.L, test.R))
    }
}