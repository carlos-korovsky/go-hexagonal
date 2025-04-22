package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type ProductInterface interface {
	IsProductValid() (bool, error)
	GetProductId() string
	GetProductName() string
	GetProductStatus() string
	GetProductPrice() float64
	EnableProduct() error
	DisableProduct() error
}

const (
	PRODUCT_STATUS_ENABLED  = "enabled"
	PRODUCT_STATUS_DISABLED = "disabled"
)

type Product struct {
	ProductId     string  `valid:"uuidv4,optional"`
	ProductName   string  `valid:"required"`
	ProductStatus string  `valid:"required,in(enabled|disabled),productStatusValidator"`
	ProductPrice  float64 `valid:"float,optional"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("productStatusValidator", func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case Product:
			if v.ProductStatus == PRODUCT_STATUS_DISABLED {
				return v.ProductPrice == 0.0
			}
			if v.ProductStatus == PRODUCT_STATUS_ENABLED {
				return v.ProductPrice > 0.0
			}
		}
		return false
	})

}

func (p *Product) GetProductId() string {
	return p.ProductId
}

func (p *Product) GetProductName() string {
	return p.ProductName
}

func (p *Product) GetProductStatus() string {
	return p.ProductStatus
}

func (p *Product) GetProductPrice() float64 {
	return p.ProductPrice
}

func NewProduct(productId string, productName string, productStatus string, productPrice float64) *Product {
	return &Product{
		ProductId:     productId,
		ProductName:   productName,
		ProductStatus: productStatus,
		ProductPrice:  productPrice,
	}
}

func (p *Product) IsProductValid() (bool, error) {

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) EnableProduct() error {
	if p.ProductPrice > 0 {
		p.ProductStatus = PRODUCT_STATUS_ENABLED
		return nil
	}
	return errors.New("product price must be greater than 0 to enable")
}

func (p *Product) DisableProduct() error {
	if p.ProductPrice == 0 {
		p.ProductStatus = PRODUCT_STATUS_DISABLED
		return nil
	}
	return errors.New("product price must be 0 to disable")
}
