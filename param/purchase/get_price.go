package param

type GetPriceRequest struct {
	ItemID uint `json:"item_id"`
}

type GetPriceResponse struct {
	Price float64 `json:"price"`
}
