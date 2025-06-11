package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/entity"
)

type UserRepository interface {
	FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
}

type userRepositoryImpl struct {
	db DBTX
}

func NewUserRepository(db DBTX) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByIdentifier(ctx context.Context, identifier string) (*entity.User, error) {
	query := `
	SELECT
		id,
		username,
		email,
		hash_password,
		role,
		created_at,
		updated_at
	FROM
		users
	WHERE
		deleted_at IS NULL
		AND (username = $1 OR email = $1)
	`

	var user entity.User
	if err := r.db.QueryRowContext(ctx, query, identifier).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.HashPassword,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
