package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
)

func NewUnauthorizedError() *ResponseError {
	msg := constant.UnauthorizedErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusUnauthorized, msg)
}
