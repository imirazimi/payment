package storage

import (
	"context"
	entity "payment/entity/ip_lookup"
)

type Config struct {
	Path string `koanf:"path"`
}

type parser interface {
	ParseCountry(ctx context.Context, path, ip string) (entity.Country, error)
}

type storage struct {
	config Config
	parser parser
}

func New(config Config, parser parser) storage {
	return storage{
		config: config,
		parser: parser,
	}
}
