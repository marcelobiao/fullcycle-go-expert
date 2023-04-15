package tax2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://github.com/stretchr/testify
func TestCalculateTax(t *testing.T) {
	result, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, result)

	result, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, result)
}
