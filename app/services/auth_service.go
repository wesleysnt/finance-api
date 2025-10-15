package services

import (
	"context"

	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/app/repositories"
	"github.com/wesleysnt/finance-api/app/responses"
	"github.com/wesleysnt/finance-api/app/schemas"
	"github.com/wesleysnt/finance-api/pkg/auth"
)

type AuthService interface {
	Login(request *requests.LoginRequest, ctx context.Context) (*responses.LoginResponse, error)
	Register(request *requests.RegisterRequest, ctx context.Context) (*responses.RegisterResponse, error)
}

type authService struct {
	jwt      auth.JWTService
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository, jwt auth.JWTService) AuthService {
	return &authService{
		jwt:      jwt,
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
			Message: "Invalid credential",
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
		Email: request.Email,
		Token: token,
	}, nil
}

func (s *authService) Register(request *requests.RegisterRequest, ctx context.Context) (*responses.RegisterResponse, error) {
	_, err := s.userRepo.GetUserByEmail(request.Email, ctx)
	if err == nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Email already in use",
		}
	}

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: &request.Password,
		Currency: request.Currency,
	}

	err = s.userRepo.CreateUser(user, ctx)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "Failed to create user",
		}
	}

	token, err := s.jwt.GenerateToken(int(user.ID), user.Email, "user")
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "Failed to generate token",
		}
	}
	return &responses.RegisterResponse{
		Email: user.Email,
		Token: token,
	}, nil

}
