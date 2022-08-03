package features

import (
	"api-shops/models"
)

type SaleFeature struct {
	sale models.SaleRepository
}

func NewSaleFeature(sale models.SaleRepository) SaleFeature {
	return SaleFeature{
		sale: sale,
	}
}

func (feat SaleFeature) CreateSale(sale models.Sale) error {

	sale.Total = feat.GetTotalSale(&sale)
	return feat.sale.CreateSale()
}

func (feat SaleFeature) GetTotalSale(sale *models.Sale) int {
	// products := NewProductFeature()

	// var total int
	// for i, item := range sale.Products{
	// 	sale[i] =
	// }

	return 1
}
