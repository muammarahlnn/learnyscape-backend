package service

import (
	"context"

	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	"github.com/muammarahlnn/user-service/internal/dto"
	"github.com/muammarahlnn/user-service/internal/entity"
	"github.com/muammarahlnn/user-service/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
}

type userServiceImpl struct {
	dataStore repository.DataStore
	hasher    encryptutil.Hasher
}

func NewUserService(
	dataStore repository.DataStore,
	hasher encryptutil.Hasher,
) UserService {
	return &userServiceImpl{
		dataStore: dataStore,
		hasher:    hasher,
	}
}

func (s *userServiceImpl) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	var res *dto.UserResponse
	err := s.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
		userRepo := ds.UserRepository()

		user, err := userRepo.FindByUsername(ctx, req.Username)
		if err != nil {
			return err
		}
		if user != nil {
		}

		user, err = userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if user != nil {
		}

		hashedPassword, err := s.hasher.Hash(req.Password)
		if err != nil {
			return err
		}

		user, err = userRepo.Create(ctx, &entity.CrateUserParams{
			Username:     req.Username,
			Email:        req.Email,
			FullName:     req.FullName,
			RoleID:       req.RoleID,
			HashPassword: hashedPassword,
		})
		if err != nil {
		}

		res = dto.ToUserResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
