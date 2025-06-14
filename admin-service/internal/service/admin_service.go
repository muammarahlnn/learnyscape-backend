package service

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/client"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/repository"
)

type AdminService interface {
	GetRoles(ctx context.Context) ([]*dto.RoleResponse, error)
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
}

type adsminServiceImpl struct {
	dataStore  repository.DataStore
	userClient client.UserClient
}

func NewAdminService(
	dataStore repository.DataStore,
	userClient client.UserClient,
) AdminService {
	return &adsminServiceImpl{
		dataStore:  dataStore,
		userClient: userClient,
	}
}

func (s *adsminServiceImpl) GetRoles(ctx context.Context) ([]*dto.RoleResponse, error) {
	roles, err := s.dataStore.RoleRepository().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToRoleResponses(roles), nil
}

func (s *adsminServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userClient.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
