package server

import (
	"api-shops/models"

	"github.com/gofiber/fiber/v2"
)

func (s Server) CreateSale(ctx *fiber.Ctx) error {
	var sale models.Sale
	err := ctx.BodyParser(&sale)

	if err != nil {
		return err
	}

	err = s.sale.CreateSale(sale)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).SendString("venta creada")
}
