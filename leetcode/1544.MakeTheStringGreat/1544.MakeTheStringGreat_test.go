package makethestringgreat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp string 
    }{
        {In:  "leEeetcode", Exp: "leetcode" },
        {In:  "abBAcC", Exp: "" },
        {In:  "s", Exp: "s" },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,makeGood(test.In))
    }
}

