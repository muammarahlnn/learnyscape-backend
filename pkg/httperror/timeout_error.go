package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
)

func NewTimeoutError() *ResponseError {
	msg := constant.RequestTimeoutErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusGatewayTimeout, msg)
}
