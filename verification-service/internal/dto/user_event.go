package dto

import "fmt"

type UserCreatedEvent struct {
	Id int64 `json:"id"`
}

func (e *UserCreatedEvent) ID() string {
	return fmt.Sprintf("%d", e.Id)
}
