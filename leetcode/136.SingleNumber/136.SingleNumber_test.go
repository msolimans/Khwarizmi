package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 3, singleNumber([]int{1,2,3,2,1}))
}