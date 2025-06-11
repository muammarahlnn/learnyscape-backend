package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/handler"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	appHandler.Route(router)

	BootstrapAuth(router)
}
