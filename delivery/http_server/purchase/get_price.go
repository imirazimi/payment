package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	param "payment/param/purchase"
	"strconv"
)

func (h Handler) GetPrice(ctx *gin.Context) {

	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	itemID64, cErr := strconv.ParseUint(id, 10, 64)
	if cErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	itemID := uint(itemID64)
	resp, cErr := h.purchaseSvc.GetPrice(ctx, param.GetPriceRequest{
		ItemID: itemID,
	})
	if cErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
