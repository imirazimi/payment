package storage

import (
	"context"
	entity "payment/entity/ip_lookup"
	richerror "payment/pkg/rich_error"
)

func (s storage) Get(ctx context.Context, ip string) (entity.Country, error) {
	const op = richerror.Op("iplookupstorage.Get")
	country, pErr := s.parser.ParseCountry(ctx, s.config.Path, ip)
	if pErr != nil {
		return entity.Country{}, richerror.New(op).WithKind(richerror.KindUnexpected)
	}
	return country, nil
}
