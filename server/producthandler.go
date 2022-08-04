package server

import "github.com/gofiber/fiber/v2"

func (s Server) GetProduct(ctx *fiber.Ctx) error {
	products, err := s.Feature.GetProducts()
	if err != nil {
		return err
	}

	return ctx.JSON(products)

}

func (s Server) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	product, err := s.Feature.GetProductByID(id)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(product)

}
