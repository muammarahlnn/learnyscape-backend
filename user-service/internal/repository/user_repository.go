package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muammarahlnn/user-service/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userRepositoryImpl struct {
	db DBTX
}

func NewUserRepository(db DBTX) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error) {
	query := `
	WITH inserted_users AS (
		INSERT INTO
			users (
				username,
				email,
				hash_password,
				full_name,
				role_id
			)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING
			id,
			username,
			email,
			full_name,
			profile_pic_url,
			is_verified,
			created_at,
			updated_at,
			role_id
	)
	SELECT
		iu.id,
		iu.username,
		iu.email,
		iu.full_name,
		iu.profile_pic_url,
		iu.is_verified,
		r.name,
		iu.created_at,
		iu.updated_at
	FROM
		inserted_users AS iu
	JOIN
		roles AS r
		ON iu.role_id = r.id
	`

	var user entity.User
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.Username,
		params.Email,
		params.HashPassword,
		params.FullName,
		params.RoleID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.ProfilePicURL,
		&user.IsVerified,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
	SELECT
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.profile_pic_url,
		u.is_verified,
		r.name,
		u.created_at,
		u.updated_at
	FROM
		users AS u
	JOIN
		roles AS r
		ON u.role_id = r.id
	WHERE
		u.deleted_at IS NULL
		AND u.username = $1
	`

	var user entity.User
	if err := r.db.QueryRowxContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.ProfilePicURL,
		&user.IsVerified,
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

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
	SELECT
		u.id,
		u.username,
		u.email,
		u.full_name,
		u.profile_pic_url,
		u.is_verified,
		r.name,
		u.created_at,
		u.updated_at
	FROM
		users AS u
	JOIN
		roles AS r
		ON u.role_id = r.id
	WHERE
		u.deleted_at IS NULL
		AND u.email = $1
	`

	var user entity.User
	if err := r.db.QueryRowxContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.ProfilePicURL,
		&user.IsVerified,
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
