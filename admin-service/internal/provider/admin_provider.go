package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/client"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/config"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/handler"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/service"
)

func BootstrapAdmin(cfg *config.Config, router *gin.Engine) {
	userClient := client.NewUserClient(cfg.Client.UserURL)

	adminService := service.NewAdminService(dataStore, userClient)
	adminHandler := handler.NewAdminHandler(adminService)

	adminHandler.Route(router)
}
