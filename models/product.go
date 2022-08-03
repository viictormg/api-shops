package models

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Price    int    `json:"price,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Subtotal int    `json:"subtotal,omitempty"`
	IDShop   string `json:"idShop"`
}
