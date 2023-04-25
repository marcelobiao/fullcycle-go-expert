package tax2

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error Save Tax"))

	err := CalculateTaxSave(1000.0, repository)
	assert.Nil(t, err)
	err = CalculateTaxSave(2000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxSave(0.0, repository)
	assert.Error(t, err, "Error Save Tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
