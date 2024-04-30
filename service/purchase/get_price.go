package service

import (
	"context"
	param "payment/param/purchase"
	taxparam "payment/param/tax"
	richerror "payment/pkg/rich_error"
)

func (s Service) GetPrice(ctx context.Context, req param.GetPriceRequest) (param.GetPriceResponse, error) {
	const op = richerror.Op("purchaseservice.Get")
	value := ctx.Value("country")
	country, ok := value.(string)
	if !ok {
		return param.GetPriceResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected)
	}
	taxSvcResp, taxSvcErr := s.taxSvc.Get(ctx, taxparam.GetTaxRequest{
		Country: country,
	})
	if taxSvcErr != nil {
		return param.GetPriceResponse{}, richerror.New(op).WithErr(taxSvcErr).WithKind(richerror.KindUnexpected)
	}
	item, gErr := s.repo.Get(ctx, req.ItemID)
	if gErr != nil {
		return param.GetPriceResponse{}, richerror.New(op).WithErr(gErr).WithKind(richerror.KindUnexpected)
	}
	price := (taxSvcResp.Tax.Percentage * item.Price / 100) + item.Price
	return param.GetPriceResponse{Price: price}, nil
}
