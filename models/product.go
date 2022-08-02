package models

type ProductRepository interface {
	GetProducts() ([]Product, error)
}

type ProductFeature interface {
	GetProducts() ([]Product, error)
}

type Product struct {
	Name string `json:"name"`
}
