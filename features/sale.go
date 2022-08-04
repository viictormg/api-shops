package features

import (
	"api-shops/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
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

	createPoint(sale.Total)

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

func createPoint(amount int) {
	var client fasthttp.Client
	reqTimeout := time.Duration(30) * time.Second

	var Response string

	requestGuide := models.RequestSavePoints{
		IdCustomer: "1d1326a8-12d5-11ed-ab31-9828a64b41f5",
		Amount:     amount,
	}

	reqEntityBytes, _ := json.Marshal(requestGuide)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://127.0.0.1:8000/api/points")

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))

	req.SetBodyRaw(reqEntityBytes)
	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)

	if err == nil {
		statusCode := resp.StatusCode()
		if statusCode == http.StatusOK {
			respBody := resp.Body()
			response := &Response
			err = json.Unmarshal(respBody, &response)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
