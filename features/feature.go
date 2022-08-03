package features

import (
	"api-shops/repository"
)

type Feature struct {
	db repository.Repository
}

func NewFeature(db repository.Repository) Feature {
	return Feature{
		db: db,
	}
}
