package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(cfg *Config) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.DBUser,
		cfg.DBPwd,
		cfg.DBIp,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", url)

	if err != nil {
		log.Panic(err)
	}

	return db, err
}
