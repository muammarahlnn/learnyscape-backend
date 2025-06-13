package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/user-service/internal/handler"
	"github.com/muammarahlnn/user-service/internal/service"
)

func BootstrapUser(router *gin.Engine) {
	userService := service.NewUserService(dataStore, bcryptHasher)
	userHandler := handler.NewUserHanlder(userService)

	userHandler.Route(router)
}
