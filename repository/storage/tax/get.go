package storage

import (
	"context"
	entity "payment/entity/tax"
	richerror "payment/pkg/rich_error"
	strings "payment/pkg/strings"
)

func (s storage) Get(ctx context.Context, country string) (entity.Tax, error) {
	const op = richerror.Op("taxstorage.Get")
	taxes, pErr := s.parser.ParseTaxes(s.config.Path)
	if pErr != nil {
		return entity.Tax{}, richerror.New(op).WithErr(pErr).WithKind(richerror.KindUnexpected)
	}
	for _, tax := range taxes {
		if strings.ContainsI(tax.Country, country) {
			return tax, nil
		}
	}
	return entity.Tax{}, richerror.New(op).WithErr(pErr).WithKind(richerror.KindUnexpected)
}
