package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/handler"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	gatewayHandler := handler.NewGatewayHandler(cfg.Service)

	appHandler.Route(router)
	gatewayHandler.Route(router)
}
