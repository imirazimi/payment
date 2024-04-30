package param

import entity "payment/entity/ip_lookup"

type IPToCountryRequest struct {
	IP string
}

type IPToCountryResponse struct {
	Country entity.Country
}
