package dto

import (
	"time"

	"github.com/muammarahlnn/user-service/internal/entity"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	RoleID   int64  `json:"role_id" binding:"required"`
	Password string `json:"password" binding:"required,password"`
}

type UserResponse struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	ProfilePicURL *string   `json:"profile_pic_url"`
	IsVerified    bool      `json:"is_verified"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		FullName:      user.Email,
		ProfilePicURL: user.ProfilePicURL,
		IsVerified:    user.IsVerified,
		Role:          user.Role,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
