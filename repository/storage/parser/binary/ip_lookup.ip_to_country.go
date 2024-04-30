package parser

import (
	"context"
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
	entity "payment/entity/ip_lookup"
)

func (p parser) ParseCountry(ctx context.Context, path, ip string) (entity.Country, error) {
	db, oErr := ip2location.OpenDB(path)
	if oErr != nil {
		return entity.Country{}, fmt.Errorf("can not open file")
	}
	defer db.Close()
	res, rErr := db.Get_all(ip)
	if rErr != nil {
		return entity.Country{}, fmt.Errorf("can not parse file")

	}
	return entity.Country{Country: res.Country_short}, nil
}
