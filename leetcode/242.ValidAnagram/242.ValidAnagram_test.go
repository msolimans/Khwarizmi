package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
    tests := []struct{
        In string
		In2 string 
        Exp bool
    }{
        {In: "anagram", In2: "nagaram", Exp: true },
        {In: "rat", In2: "car", Exp: false },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, isAnagram(test.In, test.In2))
    }
}

func TestIsAnagram2(t *testing.T) {
    tests := []struct{
        In string
		In2 string 
        Exp bool
    }{
        {In: "anagram", In2: "nagaram", Exp: true },
        {In: "rat", In2: "car", Exp: false },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, isAnagram2(test.In, test.In2))
    }
}

