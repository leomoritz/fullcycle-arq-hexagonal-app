package application_test

import (
	"testing"

	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	valid, err := product.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be ENABLED or DISABLED", err.Error())

	product.Status = ""
	valid, err = product.IsValid()
	require.Equal(t, application.DISABLED, product.Status)
	require.True(t, valid)
	require.Nil(t, err)

	product.Status = application.ENABLED
	valid, err = product.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	product.Price = -10
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Equal(t, "the price must be greater than or equal to zero", err.Error())

	product.Price = 10
	product.Name = ""
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Equal(t, "Name: non zero value required", err.Error())

	product.Name = "Product 1"
	product.ID = "1234"
	valid, err = product.IsValid()
	require.False(t, valid)
	require.Equal(t, "ID: 1234 does not validate as uuidv4", err.Error())

	product.ID = uuid.NewV4().String()
	product.Price = 0
	valid, err = product.IsValid()
	require.True(t, valid)
	require.Nil(t, err)
}

func TestProduct_Get(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	require.Equal(t, product.ID, product.GetID())
	require.Equal(t, product.Name, product.GetName())
	require.Equal(t, product.Price, product.GetPrice())
	require.Equal(t, product.Status, product.GetStatus())
}
