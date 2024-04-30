package storage

import entity "payment/entity/purchase"

type Config struct {
	Path string `koanf:"path"`
}
type parser interface {
	ParseItems(path string) ([]entity.Item, error)
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
