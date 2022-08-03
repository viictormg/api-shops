package features

import (
	"api-shops/models"
)

func (feat Feature) CreateSale(sale models.Sale) error {
	err := feat.GetPriceProducts(&sale)
	if err != nil {
		return err
	}

	id, err := feat.db.CreateSale(sale)

	if err != nil {
		return err
	}
	err = feat.db.CreateDetailSale(id, sale.Products)

	return err
}

func (feat Feature) GetPriceProducts(sale *models.Sale) error {
	var total int
	for i, item := range sale.Products {
		product, err := feat.GetProductByID(item.ID)

		if err != nil {
			return err
		}
		subtotal := item.Quantity * product.Price
		sale.Products[i] = models.Product{
			ID:       item.ID,
			Price:    product.Price,
			Quantity: item.Quantity,
			Subtotal: subtotal,
		}
		total += subtotal
	}
	sale.Total = total
	return nil
}
