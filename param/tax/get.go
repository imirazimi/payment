package param

import entity "payment/entity/tax"

type GetTaxRequest struct {
	Country string
}

type GetTaxResponse struct {
	Tax entity.Tax
}
