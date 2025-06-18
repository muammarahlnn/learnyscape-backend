package dto

import (
	"fmt"
)

type UserCreatedEvent struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (e *UserCreatedEvent) ID() string {
	return fmt.Sprintf("%d", e.Id)
}
