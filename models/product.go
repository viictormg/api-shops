package models

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductID(id string) (Product, error)
}

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Price    int    `json:"price,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
