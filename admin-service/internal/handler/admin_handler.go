package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/service"
	ginutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/gin"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) Route(r *gin.Engine) {
	r.GET("/roles", h.getRoles)
	r.POST("/users", h.createUser)
}

func (h *AdminHandler) getRoles(ctx *gin.Context) {
	res, err := h.adminService.GetRoles(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AdminHandler) createUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.adminService.CreateUser(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseCreated(ctx, res)
}
