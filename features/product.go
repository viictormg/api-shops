package features

import (
	"api-shops/models"
)

func (feat Feature) GetProducts() ([]models.Product, error) {
	return feat.db.GetProducts()
}

func (feat Feature) GetProductByID(id string) (models.Product, error) {
	return feat.db.GetProductID(id)
}
