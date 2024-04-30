package service

import (
	"context"
	entity "payment/entity/purchase"
	param "payment/param/tax"
	storage "payment/repository/storage/purchase"
)

type Config struct {
	Storage storage.Config `koanf:"storage"`
}
type repository interface {
	Get(ctx context.Context, ID uint) (entity.Item, error)
}
type taxService interface {
	Get(ctx context.Context, req param.GetTaxRequest) (param.GetTaxResponse, error)
}

type Service struct {
	config Config
	repo   repository
	taxSvc taxService
}

func New(config Config, repo repository, taxSvc taxService) Service {
	return Service{
		config: config,
		repo:   repo,
		taxSvc: taxSvc,
	}
}
