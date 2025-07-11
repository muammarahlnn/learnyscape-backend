package dto

import "fmt"

type UserCreatedEvent struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (e *UserCreatedEvent) ID() string {
	return fmt.Sprintf("%d", e.Id)
}
