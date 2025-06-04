package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
	"github.com/muammarahlnn/learnyscape-backend/pkg/dto"
	"github.com/muammarahlnn/learnyscape-backend/pkg/httperror"
	validationutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/validation"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) < 1 {
			return
		}

		err := ctx.Errors.Last()
		switch e := err.Err.(type) {
		case validator.ValidationErrors:
			handleValidationError(ctx, e)
		case *json.SyntaxError:
			handleJsonSyntaxError(ctx)
		case *json.UnmarshalTypeError:
			handleJsonUnmarshalTypeError(ctx, e)
		case *strconv.NumError:
			handleNumError(ctx, e)
		case *time.ParseError:
			handleParseTimeError(ctx, e)
		case *httperror.ResponseError:
			handleHttpResponseError(ctx, e)
		default:
			if errors.Is(e, io.EOF) {
				handleEOFError(ctx)
				return
			}

			handleInternalServerError(ctx)
		}
	}
}

func handleJsonSyntaxError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: constant.JsonSyntaxErrorMessage,
		},
	)
}

func handleJsonUnmarshalTypeError(ctx *gin.Context, err *json.UnmarshalTypeError) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: fmt.Sprintf(constant.JsonUnMarshallTypeErrorMessage, err.Field),
		},
	)
}

func handleNumError(ctx *gin.Context, err *strconv.NumError) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: fmt.Sprintf(constant.NumErrorMessage, err.Num),
		},
	)
}

func handleParseTimeError(ctx *gin.Context, err *time.ParseError) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: fmt.Sprintf(
				"please send time in format of %s, got: %s",
				constant.ConvertGoTimeLayoutToReadable(err.Layout), err.Value,
			),
		},
	)
}

func handleHttpResponseError(ctx *gin.Context, err *httperror.ResponseError) {
	ctx.AbortWithStatusJSON(
		err.Code(),
		dto.WebResponse[any]{
			Message: err.Message(),
		},
	)
}

func handleValidationError(ctx *gin.Context, err validator.ValidationErrors) {
	ve := []*dto.FieldError{}

	for _, fe := range err {
		fieldError := &dto.FieldError{
			Field:   fe.Field(),
			Message: validationutil.TagToMsg(fe),
		}
		ve = append(ve, fieldError)
	}

	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: constant.ValidationErrorMessage,
			Errors:  ve,
		},
	)
}

func handleEOFError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		dto.WebResponse[any]{
			Message: constant.EOFErrorMessage,
		},
	)
}

func handleInternalServerError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		dto.WebResponse[any]{
			Message: constant.InternalServerErrorMessage,
		},
	)
}
