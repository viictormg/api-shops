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

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RequestSavePoints struct {
	IdCustomer string `json:"idCustomer"`
	Amount     int    `json: "amount"`
}
