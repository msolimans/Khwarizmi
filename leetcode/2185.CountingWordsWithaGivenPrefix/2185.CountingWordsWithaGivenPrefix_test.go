package countingwordswithagivenprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []string 
		Prefix string 
        Exp int 
    }{
        {In: []string{"pay","attention","practice","attend"}, Prefix: "at", Exp: 2},
        {In: []string{"leetcode","win","loops","success"}, Prefix: "code", Exp: 0},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,prefixCount(test.In, test.Prefix))
    }
}

