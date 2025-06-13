package httperror

import (
	"errors"
	"net/http"

	"github.com/muammarahlnn/learnyscape-backend/pkg/httperror"
	"github.com/muammarahlnn/user-service/internal/constant"
)

func NewUserAlreadyExistsError() *httperror.ResponseError {
	msg := constant.UserAlreadyExistsErorrMesage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusConflict, msg)
}
