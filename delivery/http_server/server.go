package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	purchasehandler "payment/delivery/http_server/purchase"
	iplookupsvc "payment/service/ip_lookup"
	purchaseservice "payment/service/purchase"
)

type Config struct {
	Port uint `koanf:"port"`
}
type server struct {
	config          Config
	router          *gin.Engine
	purchaseHandler purchasehandler.Handler
}

func New(config Config, purchaseSvc purchaseservice.Service, ipLookupSvc iplookupsvc.Service) server {
	return server{
		config:          config,
		router:          gin.New(),
		purchaseHandler: purchasehandler.New(purchaseSvc, ipLookupSvc),
	}
}

func (s server) Run() error {
	s.router.GET("", healthCheck)
	s.purchaseHandler.SetRoutes(s.router)
	address := fmt.Sprintf(":%d", s.config.Port)
	rErr := s.router.Run(address)
	return rErr
}
