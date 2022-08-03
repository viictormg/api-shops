package server

import (
	"api-shops/config"
	"api-shops/features"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg     *config.Config
	product features.ProductFeature
	sale    features.SaleFeature
}

func NewServer(cfg *config.Config, product features.ProductFeature, sale features.SaleFeature) *Server {
	return &Server{
		cfg:     cfg,
		product: product,
		sale:    sale,
	}
}

func (s Server) Run() error {

	app := fiber.New()

	s.RegisterRoute(app)

	return app.Listen(s.cfg.Host)
}
