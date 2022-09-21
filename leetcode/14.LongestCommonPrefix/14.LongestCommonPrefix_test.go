package LongestCommonPrefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{"sa","sa", ""}))
	assert.Equal(t, "", longestCommonPrefix([]string{"sa",""}))
	assert.Equal(t, "", longestCommonPrefix([]string{"","sa"}))
	assert.Equal(t, "s", longestCommonPrefix([]string{"ss","sa"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"a","sa"}))
}


func TestFindMatch(t *testing.T) {
	matched := findMatch("ss","sa", 0)
	assert.Equal(t, "s", matched)
	assert.Equal(t, "s", findMatch("s","sa", 0))
	assert.Equal(t, "", findMatch("a","sa", 0))
}
