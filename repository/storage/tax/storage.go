package storage

import entity "payment/entity/tax"

type Config struct {
	Path string `koanf:"path"`
}

type parser interface {
	ParseTaxes(path string) ([]entity.Tax, error)
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
