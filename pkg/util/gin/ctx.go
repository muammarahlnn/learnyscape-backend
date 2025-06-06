package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
)

func GetUserID(ctx *gin.Context) int64 {
	return ctx.GetInt64(constant.CTX_USER_ID)
}

func GetUserRole(ctx *gin.Context) string {
	return ctx.GetString(constant.CTX_USER_ROLE)
}
