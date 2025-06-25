package repository

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, params *entity.CreateUserParams) (*entity.User, error)
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
	WITH role_lookup AS (
		SELECT
			id
		FROM
			roles
		WHERE
			deleted_at IS NULL 
			AND name = $1
	), 
	inserted_users AS (
		INSERT INTO
			users (
				id,
				username,
				email,
				full_name,
				role_id
			)
		SELECT
			$2,
			$3,
			$4,
			$5,
			rl.id
		FROM
			role_lookup AS rl
		ON CONFLICT (id) DO UPDATE SET
			username = EXCLUDED.username,
			email = EXCLUDED.email,
			full_name = EXCLUDED.full_name,
			updated_at = EXCLUDED.updated_at,
			role_id = EXCLUDED.role_id
		RETURNING
			id,
			username,
			email,
			full_name,
			profile_pic_url,
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
		r.name,
		iu.created_at,
		iu.updated_at
	FROM
		inserted_users AS iu
	JOIN
		roles AS r ON iu.role_id = r.id
	`

	var user entity.User
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.Role,
		params.ID,
		params.Username,
		params.Email,
		params.FullName,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.ProfilePicURL,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
