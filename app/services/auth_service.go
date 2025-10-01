package services

import (
	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/app/responses"
	"github.com/wesleysnt/finance-api/app/schemas"
	"github.com/wesleysnt/finance-api/pkg/auth"
)

type AuthService struct {
	jwt auth.JWTService
}

func NewAuthService() *AuthService {
	return &AuthService{
		jwt: *auth.NewJWTService(),
	}
}

func (s *AuthService) Login(request *requests.LoginRequest) (*responses.LoginResponse, error) {
	if request.Email != "test@gmail.com" && request.Password != "password" {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "invalid email or password",
		}
	}

	jwt, err := s.jwt.GenerateToken(1, request.Email, "Test")

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
	}
	resp := responses.LoginResponse{
		Email: request.Email,
		Token: jwt,
	}

	return &resp, err

}
