package server

import "github.com/gofiber/fiber/v2"

func (s Server) GetProduct(ctx *fiber.Ctx) error {

	products, err := s.product.GetProducts()

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(products)

}
