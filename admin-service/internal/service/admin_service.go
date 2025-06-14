package service

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/admin-service/internal/repository"
)

type AdminService interface {
	GetRoles(ctx context.Context) ([]*dto.RoleResponse, error)
}

type adsminServiceImpl struct {
	dataStore repository.DataStore
}

func NewAdminService(dataStore repository.DataStore) AdminService {
	return &adsminServiceImpl{
		dataStore: dataStore,
	}
}

func (s *adsminServiceImpl) GetRoles(ctx context.Context) ([]*dto.RoleResponse, error) {
	roles, err := s.dataStore.RoleRepository().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToRoleResponses(roles), nil
}