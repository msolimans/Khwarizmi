package excelsheetcolumntitle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTitle(t *testing.T) {
	assert.Equal(t, "CV", convertToTitle(100))
	assert.Equal(t, "A", convertToTitle(1))
	assert.Equal(t, "AB", convertToTitle(28))
	assert.Equal(t, "ZY", convertToTitle(701))
}