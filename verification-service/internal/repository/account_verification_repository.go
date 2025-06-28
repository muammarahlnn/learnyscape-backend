package repository

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/verification-service/internal/entity"
)

type AccountVerificationRepository interface {
	Save(ctx context.Context, params *entity.CreateAccountVerificationParams) (*entity.AccountVerification, error)
}

type accountVerificationRepositoryImpl struct {
	db DBTX
}

func NewAccountVerificationRepository(db DBTX) AccountVerificationRepository {
	return &accountVerificationRepositoryImpl{
		db: db,
	}
}

func (r accountVerificationRepositoryImpl) Save(ctx context.Context, params *entity.CreateAccountVerificationParams) (*entity.AccountVerification, error) {
	query := `
	INSERT INTO
		user_verifications (
			user_id,
			token,
			expire_at
		)
	VALUES
		($1, $2, $3)
	ON CONFLICT (user_id) DO UPDATE SET
		user_id = EXCLUDED.user_id,
		token = EXCLUDED.token,
		expire_at = EXCLUDED.expire_at,
		updated_at = NOW()
	RETURNING
		id,
		user_id,
		token,
		expire_at,
		created_at,
		updated_at
	`

	var verification entity.AccountVerification
	if err := r.db.QueryRowContext(
		ctx,
		query,
		params.UserID,
		params.Token,
		params.ExpireAt,
	).Scan(
		&verification.ID,
		&verification.UserID,
		&verification.Token,
		&verification.ExpireAt,
		&verification.CreatedAt,
		&verification.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &verification, nil
}
