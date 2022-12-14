package server

import "github.com/gofiber/fiber/v2"

func (s Server) RegisterRoute(app *fiber.App) {
	routes := app.Group("/api")

	product := routes.Group("/product")
	product.Get("/", s.GetProduct)
	product.Get("/:id", s.GetProductByID)

	sale := routes.Group("/sale")
	sale.Post("/", s.CreateSale)

}
