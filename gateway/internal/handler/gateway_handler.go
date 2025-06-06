package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/proxy"
)

type GatewayHandler struct {
	serviceConfig *config.ServiceConfig
}

func NewGatewayHandler(serviceConfig *config.ServiceConfig) *GatewayHandler {
	return &GatewayHandler{
		serviceConfig: serviceConfig,
	}
}

func (h *GatewayHandler) Route(r *gin.Engine) {
	authProxy := proxy.NewReverseProxy(h.serviceConfig.AuthURL)

	r.Any("/auth/*path", authProxy)
}
