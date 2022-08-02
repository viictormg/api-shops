package features

import (
	"api-shops/models"
)

type ProductFeature struct {
	product models.ProductRepository
}

func NewProductFeature(product models.ProductRepository) ProductFeature {
	return ProductFeature{
		product: product,
	}
}

func (feat ProductFeature) GetProducts() ([]models.Product, error) {
	return feat.product.GetProducts()
}
