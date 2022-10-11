package uniqueemailaddresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []string 
        Exp int 
    }{
        {In: []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, Exp: 2},
        {In: []string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"}, Exp: 3},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,numUniqueEmails(test.In))
    }
}

