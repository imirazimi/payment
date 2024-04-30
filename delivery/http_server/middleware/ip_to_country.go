package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	param "payment/param/ip_lookup"
	service "payment/service/ip_lookup"
)

func IPToCountry(ipLookupSvc service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, cErr := ipLookupSvc.IPToCountry(ctx, param.IPToCountryRequest{
			IP: ctx.ClientIP(),
		})
		if cErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		}
		ctx.Set("country", resp.Country.Country)
		ctx.Next()
	}
}
