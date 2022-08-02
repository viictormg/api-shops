package repository

import (
	"api-shops/models"
	"database/sql"
)

func NewProductRepository(db *sql.DB) models.ProductRepository {
	return Repository{
		db: db,
	}
}

func (r Repository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	query := `SELECT name FROM products`

	rows, err := r.db.Query(query)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(
			&product.Name,
		)

		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, err
}
