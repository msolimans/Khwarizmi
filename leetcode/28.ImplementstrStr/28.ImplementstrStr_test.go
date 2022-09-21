package implementstrstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T){
	assert.Equal(t, -1, strStr("a", "s"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 0, strStr("", ""))
	assert.Equal(t, 0, strStr("ss", ""))
	assert.Equal(t, 2, strStr("aallo", "ll"))
	assert.Equal(t, 4, strStr("aallo", "o"))
	assert.Equal(t, 0, strStr("aallo", "a"))
}