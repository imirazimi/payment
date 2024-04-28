package main

import (
	"context"
	"fmt"
	"log"
	param "payment/param/tax"
	storage "payment/repository/storage/tax"
	service "payment/service/tax"
	"payment/storage/parser/csv"
)

func main() {
	csvParser := parser.New()
	repo := storage.New(storage.NewConfig("./storage/tax/tax.csv"), csvParser)
	svc := service.New(repo)
	ctx := context.Background()
	resp, gErr := svc.Get(ctx, param.GetTaxRequest{Country: "germany"})
	if gErr != nil {
		log.Fatalln(gErr)
	}
	fmt.Println(resp.Tax)
}
