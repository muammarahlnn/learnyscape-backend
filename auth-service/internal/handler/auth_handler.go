package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/service"
	ginutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Route(r *gin.Engine) {
	r.POST("/login", h.login)
}

func (h *AuthHandler) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Login(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}
