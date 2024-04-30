package handler

import (
	iplookupsvc "payment/service/ip_lookup"
	purchaseservice "payment/service/purchase"
)

type Handler struct {
	purchaseSvc purchaseservice.Service
	ipLookupSvc iplookupsvc.Service
}

func New(purchaseSvc purchaseservice.Service, ipLookupSvc iplookupsvc.Service) Handler {
	return Handler{
		purchaseSvc: purchaseSvc,
		ipLookupSvc: ipLookupSvc,
	}
}
