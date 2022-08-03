package server

import (
	"api-shops/models"

	"github.com/gofiber/fiber/v2"
)

func (s Server) CreateSale(ctx *fiber.Ctx) error {
	var sale models.Sale
	err := ctx.BodyParser(&sale)

	sale.IDShop = "d2683200-12d0-11ed-ab31-9828a64b41f5"

	if err != nil {
		return err
	}

	err = s.Feature.CreateSale(sale)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).SendString("venta creada")
}
