package models

type SaleRepository interface {
	CreateSale() error
}

type Sale struct {
	IDCustomer string    `json:"idCustomer"`
	Products   []Product `json:"products"`
	Total      int
	IDShop     string `json:"idShop"`
}
