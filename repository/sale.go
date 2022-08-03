package repository

import (
	"api-shops/models"
	"database/sql"
)

func NewSaleRepository(db *sql.DB) models.SaleRepository {
	return Repository{
		db: db,
	}
}

func (r Repository) CreateSale() error {
	return nil
}
