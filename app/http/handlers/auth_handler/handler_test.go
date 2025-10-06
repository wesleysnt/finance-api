package authhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/app/responses"
	"github.com/wesleysnt/finance-api/pkg"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(request *requests.LoginRequest, ctx context.Context) (*responses.LoginResponse, error) {
	args := m.Called(request, ctx)
	return args.Get(0).(*responses.LoginResponse), args.Error(1)
}

func (m *MockAuthService) Register(request *requests.RegisterRequest, ctx context.Context) (*responses.RegisterResponse, error) {
	args := m.Called(request, ctx)
	return args.Get(0).(*responses.RegisterResponse), args.Error(1)
}

func TestAuthHandler_Login(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)
	ctx := context.Background()

	request := requests.LoginRequest{
		Email:    "test@gmail.com",
		Password: "testPassword",
	}

	expectedResp := responses.LoginResponse{
		Email: "test@gmail.com",
		Token: "login-token",
	}

	service.On("Login", mock.AnythingOfType("*requests.LoginRequest"), ctx).Return(&expectedResp, nil)

	reqJson, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	service.AssertExpectations(t)
}
