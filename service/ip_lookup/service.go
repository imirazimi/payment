package service

import (
	"context"
	entity "payment/entity/ip_lookup"
	storage "payment/repository/storage/ip_lookup"
)

type Config struct {
	Storage storage.Config `koanf:"storage"`
}
type repository interface {
	Get(ctx context.Context, ip string) (entity.Country, error)
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
