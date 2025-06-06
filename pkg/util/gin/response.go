package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/pkg/dto"
)

func ResponseOK[T any](ctx *gin.Context, data T) {
	ResponseJSON(ctx, http.StatusOK, data, nil)
}

func ResponseCreated[T any](ctx *gin.Context, data T) {
	ResponseJSON(ctx, http.StatusCreated, data, nil)
}

func ResponsePagination[T any](ctx *gin.Context, data T, paging *dto.PageMetaData) {
	ResponseJSON(ctx, http.StatusOK, data, paging)
}

func ResponseNoContent(ctx *gin.Context) {
	ResponseJSON[any](ctx, http.StatusNoContent, nil, nil)
}

func ResponseJSON[T any](ctx *gin.Context, statusCode int, data T, paging *dto.PageMetaData) {
	ctx.JSON(statusCode, dto.WebResponse[T]{Data: data, Paging: paging})
}

func ResponseMessage(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, dto.WebResponse[any]{Message: message})
}
