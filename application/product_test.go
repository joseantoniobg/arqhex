package application_test

import (
	"github.com/joseantoniobg/arqhex/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		ID:     "abc",
		Name:   "Product 1",
		Status: application.DISABLED,
		Price:  10.0,
	}

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		ID:     "abc",
		Name:   "Product 1",
		Status: application.ENABLED,
		Price:  0.0,
	}

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be 0 in order to disable product", err.Error())
}
