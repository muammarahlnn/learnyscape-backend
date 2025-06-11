package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/handler"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/service"
)

func BootstrapAuth(router *gin.Engine) {
	authService := service.NewAuthService(dataStore, bcryptHasher, jwtUtil)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
