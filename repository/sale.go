package repository

import (
	"api-shops/models"

	"github.com/google/uuid"
)

func (r Repository) CreateSale(sale models.Sale) (string, error) {

	id := uuid.NewString()
	query := `INSERT INTO sales (id, total, IDCustomer, IDShop) 
				VALUES(?, ?, ?, ?)
				`
	_, err := r.db.Query(
		query,
		id,
		sale.Total,
		sale.IDCustomer,
		sale.IDShop,
	)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (r Repository) CreateDetailSale(id string, products []models.Product) error {
	for _, detail := range products {

		query := `INSERT INTO sale_detail (id, IDproduct, subtotal, quantity, IDSale) 
				VALUES(UUID(), ?, ?, ?, ?)
				`
		_, err := r.db.Query(
			query,
			detail.ID,
			detail.Subtotal,
			detail.Quantity,
			id,
		)

		if err != nil {
			return err
		}

	}

	return nil
}
