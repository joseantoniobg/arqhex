package application_test

import (
	"github.com/joseantoniobg/arqhex/application"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10.0,
	}

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()

	require.Equal(t, "invalid product status", err.Error())

	product.Status = application.ENABLED

	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1

	_, err = product.IsValid()
	require.Equal(t, "price must be greater than or equal to 0", err.Error())
}
