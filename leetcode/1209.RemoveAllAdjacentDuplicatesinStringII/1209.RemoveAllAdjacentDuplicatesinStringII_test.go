package removealladjacentduplicatesinstringii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
    tests := []struct{
        In string 
		K int 
        Exp string  
    }{
        {In: "abcd", K: 2, Exp: "abcd"},
        {In: "deeedbbcccbdaa", K: 3, Exp: "aa"},
        {In: "pbbcggttciiippooaais", K: 2, Exp: "ps"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, removeDuplicates(test.In, test.K))
    }
}