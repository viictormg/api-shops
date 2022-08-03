package repository

import (
	"api-shops/models"
)

func (r Repository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	query := `SELECT id, name, price FROM products`

	rows, err := r.db.Query(query)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, err
}

func (r Repository) GetProductID(id string) (models.Product, error) {
	var product models.Product

	query := `SELECT id, name, price FROM products WHERE id = ? `

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	return product, err
}
