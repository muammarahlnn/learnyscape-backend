package handler

import (
	"github.com/gin-gonic/gin"
	ginutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/gin"
	"github.com/muammarahlnn/user-service/internal/dto"
	"github.com/muammarahlnn/user-service/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHanlder(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Route(r *gin.Engine) {
	r.POST("", h.create)
}

func (h *UserHandler) create(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.userService.Create(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseCreated(ctx, res)
}
