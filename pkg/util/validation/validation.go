package validationutil

import (
	"github.com/go-playground/validator/v10"
	validationtype "github.com/muammarahlnn/learnyscape-backend/pkg/util/validation/type"
)

func Password(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	password := validationtype.NewPassword(data)
	return password.Validate()
}
