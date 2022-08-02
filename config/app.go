package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	DBUser string
	DBPwd  string
	DBIp   string
	DBName string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return &Config{}, err
	}

	return &Config{
		Host:   os.Getenv("HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPwd:  os.Getenv("DB_PWD"),
		DBIp:   os.Getenv("DB_IP"),
		DBName: os.Getenv("DB_NAME"),
	}, nil
}
