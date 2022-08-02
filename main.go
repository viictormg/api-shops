package main

import (
	"api-shops/config"
	"api-shops/features"
	"api-shops/repository"
	"api-shops/server"

	"log"
)

func run() error {
	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	db, err := config.InitDB(cfg)

	if err != nil {
		return err
	}

	productRepository := repository.NewProductRepository(db)
	productFeature := features.NewProductFeature(productRepository)

	srv := server.NewServer(cfg, productFeature)

	return srv.Run()

}
func main() {

	err := run()

	if err != nil {
		log.Panic(err)
	}
	// aguapanelaconlimon

}
