package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/config"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	BootstrapApp(router)
	BootstrapAdmin(cfg, router)
}
