package service

import (
	"context"
	entity "payment/entity/tax"
	storage "payment/repository/storage/tax"
)

type Config struct {
	Storage storage.Config `koanf:"storage"`
}
type repository interface {
	Get(ctx context.Context, country string) (entity.Tax, error)
}

type Service struct {
	config Config
	repo   repository
}

func New(config Config, repo repository) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}
