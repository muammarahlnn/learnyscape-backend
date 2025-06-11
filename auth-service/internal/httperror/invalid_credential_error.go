package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/pkg/httperror"
)

func NewInvalidCredentialError() *httperror.ResponseError {
	msg := constant.InvalidCredentialErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusUnauthorized, msg)
}
