package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/pkg/dto"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) Route(r *gin.Engine) {
	r.NoRoute(h.RouteNotFound)
	r.NoMethod(h.MethodNotAllowed)
	r.GET("/ping", h.Ping)
}

func (h *AppHandler) Ping(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "pong",
		},
	)
}

func (h *AppHandler) RouteNotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusNotFound,
		dto.WebResponse[any]{
			Message: "route not found",
		},
	)
}

func (h *AppHandler) MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(
		http.StatusMethodNotAllowed,
		dto.WebResponse[any]{
			Message: "method not allowed",
		},
	)
}
