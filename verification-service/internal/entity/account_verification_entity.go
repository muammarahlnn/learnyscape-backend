package entity

import "time"

type AccountVerification struct {
	ID        int64
	UserID    int64
	Token     string
	ExpireAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateAccountVerificationParams struct {
	UserID   int64
	Token    string
	ExpireAt time.Time
}
