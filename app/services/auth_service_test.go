package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/pkg/auth"
	"gorm.io/gorm"
)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(userID int, email string, role string) (string, error) {
	args := m.Called(userID, email, role)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) GenerateRefreshToken(userID int) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(token string) (*auth.Claims, error) {
	args := m.Called(token)
	return args.Get(0).(*auth.Claims), args.Error(1)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByEmail(email string, ctx context.Context) (*models.User, error) {
	args := m.Called(email, ctx)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *models.User, ctx context.Context) error {
	args := m.Called(user, ctx)
	return args.Error(0)
}

func TestLogin(t *testing.T) {
	repo := new(MockUserRepository)
	jwt := new(MockJWTService)

	service := NewAuthService(repo, jwt)

	ctx := context.Background()

	expectedPassword, err := auth.HashPassword("testPassword")

	assert.NoError(t, err)

	expectedUser := &models.User{
		Model:    gorm.Model{ID: 1},
		Email:    "testuser@example.com",
		Password: &expectedPassword, // Assume this is a hashed password
		Currency: "USD",
	}

	repo.On("GetUserByEmail", "testuser@example.com", ctx).Return(expectedUser, nil)
	jwt.On("GenerateToken", int(expectedUser.ID), expectedUser.Email, "user").Return("mocked_jwt_token", nil)

	request := &requests.LoginRequest{
		Email:    "testuser@example.com",
		Password: "testPassword",
	}

	response, err := service.Login(request, ctx)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Token)

	repo.AssertExpectations(t)
}
