package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct{
		In string 
		Exp int 
	} {
		{In: "abccccdd", Exp: 7},
		{In: "abCcccdd", Exp: 5},
		{In: "a", Exp: 1},
	}
	for _, test := range tests {
		assert.Equal(t, test.Exp, longestPalindrome(test.In))
	}
}