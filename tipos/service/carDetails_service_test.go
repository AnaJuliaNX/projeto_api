package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetDetails(t *testing.T) {
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.ID_carro)
	assert.Equal(t, "Mustang", carDetails.Carro)
	assert.Equal(t, "Ford", carDetails.Fabricante)

}
