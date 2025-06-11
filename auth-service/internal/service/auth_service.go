package service

import (
	"context"

	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/dto"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/httperror"
	"github.com/muammarahlnn/learnyscape-backend/auth-service/internal/repository"
	encryptutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/encrypt"
	jwtutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/jwt"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authServiceImpl struct {
	dataStore repository.DataStore
	hasher    encryptutil.Hasher
	jwt       jwtutil.JWTUtil
}

func NewAuthService(
	dataStore repository.DataStore,
	hasher encryptutil.Hasher,
	jwt jwtutil.JWTUtil,
) AuthService {
	return &authServiceImpl{
		dataStore: dataStore,
		hasher:    hasher,
		jwt:       jwt,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.dataStore.UserRepository().FindByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, httperror.NewInvalidCredentialError()
	}

	if !s.hasher.Check(req.Password, user.HashPassword) {
		return nil, httperror.NewInvalidCredentialError()
	}

	jwtPayload := &jwtutil.JWTPayload{
		UserID: user.ID,
		Role:   user.Role,
	}
	accessToken, err := s.jwt.SignAccess(jwtPayload)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwt.SignRefresh(jwtPayload)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
