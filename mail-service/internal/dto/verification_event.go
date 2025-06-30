package dto

import "github.com/muammarahlnn/learnyscape-backend/mail-service/internal/constant"

type SendVerificationEvent struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Token    string `json:"token"`
}

func (e *SendVerificationEvent) Key() string {
	return constant.SendVerificationKey
}
