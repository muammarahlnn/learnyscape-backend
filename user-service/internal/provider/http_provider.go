package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/user-service/internal/handler"
)

func BootstrapHttp(router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	appHandler.Route(router)

	BootstrapUser(router)
}
