package dto

import (
	"fmt"

	"github.com/muammarahlnn/user-service/internal/entity"
)

type UserCreatedEvent struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	HashPassword string `json:"hash_password"`
	Role         string `json:"role"`
}

func (e *UserCreatedEvent) ID() string {
	return fmt.Sprintf("%d", e.Id)
}

func ToUserCreatedEvent(user *entity.User) *UserCreatedEvent {
	return &UserCreatedEvent{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		FullName:     user.FullName,
		HashPassword: user.HashPassword,
		Role:         user.Role,
	}
}
