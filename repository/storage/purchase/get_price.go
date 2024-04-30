package storage

import (
	"context"
	entity "payment/entity/purchase"
	richerror "payment/pkg/rich_error"
)

func (s storage) Get(ctx context.Context, id uint) (entity.Item, error) {
	const op = richerror.Op("purchasestorage.Get")
	items, pErr := s.parser.ParseItems(s.config.Path)
	if pErr != nil {
		return entity.Item{}, richerror.New(op).WithErr(pErr).WithKind(richerror.KindUnexpected)
	}
	for _, item := range items {
		if id == item.ID {
			return item, nil
		}
	}
	return entity.Item{}, richerror.New(op).WithErr(pErr).WithKind(richerror.KindUnexpected)
}
