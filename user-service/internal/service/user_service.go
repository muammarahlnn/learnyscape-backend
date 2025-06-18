package service

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	"github.com/muammarahlnn/user-service/internal/dto"
	"github.com/muammarahlnn/user-service/internal/entity"
	"github.com/muammarahlnn/user-service/internal/httperror"
	"github.com/muammarahlnn/user-service/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
}

type userServiceImpl struct {
	dataStore           repository.DataStore
	hasher              encryptutil.Hasher
	userCreatedProducer mq.KafkaProducer
}

func NewUserService(
	dataStore repository.DataStore,
	hasher encryptutil.Hasher,
	userCreatedProducer mq.KafkaProducer,
) UserService {
	return &userServiceImpl{
		dataStore:           dataStore,
		hasher:              hasher,
		userCreatedProducer: userCreatedProducer,
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
			return httperror.NewUserAlreadyExistsError()
		}

		user, err = userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if user != nil {
			return httperror.NewUserAlreadyExistsError()
		}

		hashedPassword, err := s.hasher.Hash(req.Password)
		if err != nil {
			return err
		}

		user, err = userRepo.Create(ctx, &entity.CreateUserParams{
			Username:     req.Username,
			Email:        req.Email,
			FullName:     req.FullName,
			RoleID:       req.RoleID,
			HashPassword: hashedPassword,
		})
		if err != nil {
			return err
		}

		if err := s.userCreatedProducer.Send(ctx, dto.ToUserCreatedEvent(user)); err != nil {
			return err
		}

		res = dto.ToUserResponse(user)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
