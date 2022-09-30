package crawlerlogfolder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
    tests := []struct{
        In []string 
        Exp int 
    }{
        {In: []string{"d1/","d2/","../","d21/","./"}, Exp: 2},
        {In: []string{"d1/","d2/","./","d3/","../","d31/"}, Exp: 3},
        {In: []string{"d1/","../","../","../"}, Exp: 0},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, minOperations(test.In))
    }
}
 