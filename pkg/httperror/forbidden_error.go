package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
)

func NewForbiddenError() *ResponseError {
	msg := constant.ForbiddenErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusForbidden, msg)
}
