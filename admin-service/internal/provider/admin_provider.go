package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/handler"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/service"
)

func BootstrapAdmin(router *gin.Engine) {
	adminService := service.NewAdminService(dataStore)
	adminHandler := handler.NewAdminHandler(adminService)

	adminHandler.Route(router)
}
