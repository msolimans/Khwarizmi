package kthmissingpositivenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestFindKthPositive(t *testing.T) {
	assert.Equal(t, 9, findKthPositive([]int{2,3,4,7,11}, 5))
	assert.Equal(t, 6, findKthPositive([]int{1,2,3,4}, 2))
	assert.Equal(t, 3, findKthPositive([]int{1}, 2))
}