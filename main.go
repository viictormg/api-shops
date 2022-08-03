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

	Repository := repository.NewRepository(db)

	Feature := features.NewFeature(Repository)

	srv := server.NewServer(cfg, Feature)

	return srv.Run()

}
func main() {

	err := run()

	if err != nil {
		log.Panic(err)
	}
	// aguapanelaconlimon

}
