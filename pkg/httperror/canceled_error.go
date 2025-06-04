package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
)

func NewCanceledError() *ResponseError {
	msg := constant.RequestCanceledErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusRequestTimeout, msg)
}
