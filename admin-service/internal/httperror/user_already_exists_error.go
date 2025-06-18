package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/constant"
	"github.com/muammarahlnn/learnyscape-backend/pkg/httperror"
)

func NewUserAlreadyExistsError() *httperror.ResponseError {
	msg := constant.UserAlreadyExistsErorrMesage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusConflict, msg)
}
