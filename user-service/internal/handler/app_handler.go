package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) Route(r *gin.Engine) {
	r.GET("/ping", h.Ping)
	r.NoRoute(h.RouteNotFound)
	r.NoMethod(h.MethodNotAllowed)
}

func (h *AppHandler) Ping(ctx *gin.Context) {
	ginutil.ResponseMessage(ctx, http.StatusOK, "pong")
}

func (h *AppHandler) RouteNotFound(ctx *gin.Context) {
	ginutil.ResponseMessage(ctx, http.StatusNotFound, "route not found")
}

func (h *AppHandler) MethodNotAllowed(ctx *gin.Context) {
	ginutil.ResponseMessage(ctx, http.StatusMethodNotAllowed, "method not allowed")
}
