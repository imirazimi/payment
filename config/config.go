package config

import (
	httpserver "payment/delivery/http_server"
	iplookupstorage "payment/repository/storage/ip_lookup"
	purchasestorage "payment/repository/storage/purchase"
	taxstorage "payment/repository/storage/tax"
	iplookupservice "payment/service/ip_lookup"
	purchaseservice "payment/service/purchase"
	taxservice "payment/service/tax"
)

type Config struct {
	TaxSvc      taxservice.Config      `koanf:"tax_service"`
	PurchaseSvc purchaseservice.Config `koanf:"purchase_service"`
	IPLookupSvc iplookupservice.Config `koanf:"ip_lookup_service"`
	HttpServer  httpserver.Config      `koanf:"http_server"`
}

const (
	taxStoragePath  = "storage/tax.tax.csv"
	itemStoragePath = "storage/purchase.item.csv"
	IPCountryPath   = "storage/ip_lookup.ip_country.bin"
	httpServerPort  = 8585
)

func defaultConfig() Config {
	return Config{
		TaxSvc: taxservice.Config{
			Storage: taxstorage.Config{
				Path: taxStoragePath,
			},
		},
		PurchaseSvc: purchaseservice.Config{
			Storage: purchasestorage.Config{
				Path: itemStoragePath,
			},
		},
		IPLookupSvc: iplookupservice.Config{
			Storage: iplookupstorage.Config{
				Path: IPCountryPath,
			},
		},
		HttpServer: httpserver.Config{
			Port: httpServerPort,
		},
	}
}
