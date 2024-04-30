package service

import (
	"context"
	param "payment/param/tax"
	richerror "payment/pkg/rich_error"
)

func (s Service) Get(ctx context.Context, req param.GetTaxRequest) (param.GetTaxResponse, error) {
	const op = richerror.Op("taxservice.Get")
	tax, gErr := s.repo.Get(ctx, req.Country)
	if gErr != nil {
		return param.GetTaxResponse{}, richerror.New(op).WithErr(gErr).WithKind(richerror.KindUnexpected)
	}
	return param.GetTaxResponse{Tax: tax}, nil
}
