package services

import (
	"context"

	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/app/repositories"
	"github.com/wesleysnt/finance-api/app/responses"
	"github.com/wesleysnt/finance-api/app/schemas"
	"github.com/wesleysnt/finance-api/pkg/auth"
)

type AuthService interface {
	Login(request *requests.LoginRequest, ctx context.Context) (*responses.LoginResponse, error)
}

type authService struct {
	jwt      auth.JWTService
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		jwt:      *auth.NewJWTService(),
		userRepo: userRepo,
	}
}

func (s *authService) Login(request *requests.LoginRequest, ctx context.Context) (*responses.LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(request.Email, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "User not found",
		}
	}

	if !auth.ComparePassword(request.Password, *user.Password) {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnauthorized,
			Message: "Invalid password",
		}
	}

	token, err := s.jwt.GenerateToken(int(user.ID), user.Email, "user")
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "Failed to generate token",
		}
	}

	return &responses.LoginResponse{
		Token: token,
	}, nil
}
