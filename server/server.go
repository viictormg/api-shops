package server

import (
	"api-shops/config"
	"api-shops/features"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg     *config.Config
	Feature features.Feature
}

func NewServer(cfg *config.Config, Feature features.Feature) *Server {
	return &Server{
		cfg:     cfg,
		Feature: Feature,
	}
}

func (s Server) Run() error {

	app := fiber.New()

	s.RegisterRoute(app)

	return app.Listen(s.cfg.Host)
}
