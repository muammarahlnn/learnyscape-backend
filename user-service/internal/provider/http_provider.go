package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/user-service/internal/config"
	"github.com/muammarahlnn/user-service/internal/handler"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	appHandler.Route(router)

	BootstrapUser(cfg, router)
}
