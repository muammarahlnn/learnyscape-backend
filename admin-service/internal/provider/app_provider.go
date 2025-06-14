package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/handler"
)

func BootstrapApp(router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	appHandler.Route(router)
}
