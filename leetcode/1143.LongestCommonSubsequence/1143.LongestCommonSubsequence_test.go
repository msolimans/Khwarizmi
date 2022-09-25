package longestcommonsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubSequence(t *testing.T) {
	assert.Equal(t, 3, longestCommonSubsequence("abcde","ace"))
}
