package application_test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/carlos-korovsky/go-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

func TestIsProductValid(t *testing.T) {
	tests := []struct {
		name      string
		product   *application.Product
		expectErr bool
	}{
		{
			name: "valid product",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_ENABLED,
				100.0,
			),
			expectErr: false,
		},
		{
			name: "valid product - no id",
			product: application.NewProduct(
				"",
				"Test Product",
				application.PRODUCT_STATUS_ENABLED,
				100.0,
			),
			expectErr: false,
		},
		{
			name: "valid product - disabled",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_DISABLED,
				0.0,
			),
			expectErr: false,
		},
		{
			name: "invalid product id",
			product: application.NewProduct(
				"123e4567-e89b-12d3-a456-426614174000",
				"Test Product",
				application.PRODUCT_STATUS_ENABLED,
				100.0,
			),
			expectErr: true,
		},
		{
			name: "invalid product with empty name",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"",
				application.PRODUCT_STATUS_ENABLED,
				100.0,
			),
			expectErr: true,
		},
		{
			name: "invalid product with invalid status",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				"invalid_status",
				100.0,
			),
			expectErr: true,
		},
		{
			name: "invalid product with status enabled and zero price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"test",
				application.PRODUCT_STATUS_ENABLED,
				0.0,
			),
			expectErr: true,
		},
		{
			name: "invalid product with status disabled and non-zero price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"test",
				application.PRODUCT_STATUS_DISABLED,
				100.0,
			),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Running test: ", tt.name)
			valid, err := tt.product.IsProductValid()
			if err != nil {
				errs := err.(govalidator.Errors).Errors()
				for _, e := range errs {
					fmt.Println(e.Error())
				}
			}
			if tt.expectErr {
				assert.False(t, valid)
				assert.Error(t, err)
			} else {
				assert.True(t, valid)
				assert.NoError(t, err)
			}
		})
	}
}

func TestEnableProduct(t *testing.T) {
	tests := []struct {
		name      string
		product   *application.Product
		expectErr bool
	}{
		{
			name: "enable product with valid price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_DISABLED,
				100.0,
			),
			expectErr: false,
		},
		{
			name: "fail to enable product with zero price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_DISABLED,
				0.0,
			),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.EnableProduct()
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, application.PRODUCT_STATUS_ENABLED, tt.product.GetProductStatus())
			}
		})
	}
}

func TestDisableProduct(t *testing.T) {
	tests := []struct {
		name      string
		product   *application.Product
		expectErr bool
	}{
		{
			name: "disable product with zero price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_ENABLED,
				0.0,
			),
			expectErr: false,
		},
		{
			name: "fail to disable product with non-zero price",
			product: application.NewProduct(
				"99133bea-f29c-4e9a-ac81-198645f649ce",
				"Test Product",
				application.PRODUCT_STATUS_ENABLED,
				100.0,
			),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.DisableProduct()
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, application.PRODUCT_STATUS_DISABLED, tt.product.GetProductStatus())
			}
		})
	}
}
