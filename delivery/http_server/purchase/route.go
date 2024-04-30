package handler

import (
	"github.com/gin-gonic/gin"
	"payment/delivery/http_server/middleware"
)

func (h Handler) SetRoutes(router *gin.Engine) {
	r := router.Group("/purchases")
	r.GET("/items/:id", middleware.IPToCountry(h.ipLookupSvc), h.GetPrice)
}
