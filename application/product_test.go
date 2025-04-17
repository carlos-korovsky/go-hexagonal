package application_test

import (
	"testing"

	"github.com/carlos-korovsky/go-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

func TestApplicationProduct_EnableProduct(t *testing.T) {
	product := application.NewProduct("1", "Test Product", application.PRODUCT_STATUS_DISABLED, 10.0)

	err := product.EnableProduct()
	assert.Nil(t, err)
	assert.Equal(t, application.PRODUCT_STATUS_ENABLED, product.GetProductStatus())
}

func TestApplicationProduct_EnableProductError(t *testing.T) {
	product := application.NewProduct("1", "Test Product", application.PRODUCT_STATUS_DISABLED, 0.0)

	err := product.EnableProduct()
	assert.NotNil(t, err)
	assert.Equal(t, "product price must be greater than 0 to enable", err.Error())
	assert.Equal(t, application.PRODUCT_STATUS_DISABLED, product.GetProductStatus())
}

func TestApplicationProduct_DisableProduct(t *testing.T) {
	product := application.NewProduct("1", "Test Product", application.PRODUCT_STATUS_ENABLED, 0.0)

	err := product.DisableProduct()
	assert.Nil(t, err)
	assert.Equal(t, application.PRODUCT_STATUS_DISABLED, product.GetProductStatus())
}

func TestApplicationProduct_DisableProductError(t *testing.T) {
	product := application.NewProduct("1", "Test Product", application.PRODUCT_STATUS_ENABLED, 10.0)

	err := product.DisableProduct()
	assert.NotNil(t, err)
	assert.Equal(t, "product price must be 0 to disable", err.Error())
	assert.Equal(t, application.PRODUCT_STATUS_ENABLED, product.GetProductStatus())
}

func TestApplicationProduct_IsProductValid(t *testing.T) {
	product := application.NewProduct("1", "Test Product", application.PRODUCT_STATUS_ENABLED, 10.0)

	valid, err := product.IsProductValid()
	assert.True(t, valid)
	assert.Nil(t, err)
}
