package entity

import "time"

type User struct {
	ID            int64
	Username      string
	Email         string
	FullName      string
	HashPassword  string
	ProfilePicURL *string
	IsVerified    bool
	Role          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateUserParams struct {
	Username     string
	Email        string
	FullName     string
	RoleID       int64
	HashPassword string
}
