package service

import (
	"context"
	"payment/param/ip_lookup"
	richerror "payment/pkg/rich_error"
)

func (s Service) IPToCountry(ctx context.Context, req param.IPToCountryRequest) (param.IPToCountryResponse, error) {
	const op = richerror.Op("iplookupservice.IPToCountry")
	country, gErr := s.repo.Get(ctx, req.IP)
	if gErr != nil {
		return param.IPToCountryResponse{}, richerror.New(op).WithErr(gErr).WithKind(richerror.KindUnexpected)
	}
	return param.IPToCountryResponse{Country: country}, nil
}
