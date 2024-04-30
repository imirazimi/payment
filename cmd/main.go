package main

import (
	"log"
	"payment/config"
	"payment/delivery/http_server"
	iplookupstorage "payment/repository/storage/ip_lookup"
	binaryparser "payment/repository/storage/parser/binary"
	csvparser "payment/repository/storage/parser/csv"
	purchasestorage "payment/repository/storage/purchase"
	taxstorage "payment/repository/storage/tax"
	iplookupsvc "payment/service/ip_lookup"
	purchaseservice "payment/service/purchase"
	taxservice "payment/service/tax"
)

func main() {

	config := config.Get()
	purchaseSvc, ipLookupSvc := setupServices(config)
	server := httpserver.New(config.HttpServer, purchaseSvc, ipLookupSvc)
	if rErr := server.Run(); rErr != nil {
		log.Fatalln(rErr)
	}

}

func setupServices(config config.Config) (purchaseservice.Service, iplookupsvc.Service) {
	ipLookupSvc := iplookupsvc.New(
		config.IPLookupSvc,
		iplookupstorage.New(config.IPLookupSvc.Storage, binaryparser.New()),
	)
	purchaseSvc := purchaseservice.New(
		config.PurchaseSvc,
		purchasestorage.New(config.PurchaseSvc.Storage, csvparser.New()),
		taxservice.New(config.TaxSvc, taxstorage.New(config.TaxSvc.Storage, csvparser.New())),
	)
	return purchaseSvc, ipLookupSvc
}
