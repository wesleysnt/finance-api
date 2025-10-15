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
	"github.com/wesleysnt/finance-api/app/schemas"
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

// FAIL TEST 1: Invalid JSON
func TestAuthHandler_Login_InvalidJSON(t *testing.T) {
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Login(c)

	assert.NoError(t, err) // Handler handles error internally
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Service should NOT be called
	mockService.AssertNotCalled(t, "Login")
}

func TestAuthHandler_Login_ValidationError(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)

	request := requests.LoginRequest{
		Email:    "",
		Password: "password",
	}

	reqJson, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Nil(t, response["data"])
	assert.Contains(t, response["message"], "required")

}

func TestAuthHandler_Login_LoginFailed(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)

	requestBody := requests.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	ctx := context.Background()
	service.On("Login", mock.AnythingOfType("*requests.LoginRequest"), ctx).Return(&responses.LoginResponse{}, &schemas.ResponseApiError{Status: schemas.ApiErrorUnauthorized, Message: "error"})

	reqJson, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	errorHand := handler.Login(c)
	assert.NoError(t, errorHand)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Nil(t, response["data"])
}

func TestAuthHandler_Register(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)
	ctx := context.Background()

	service.On("Register", mock.AnythingOfType("*requests.RegisterRequest"), ctx).Return(
		&responses.RegisterResponse{
			Email: "newUser@mail.com",
			Token: "mock_jwt_token",
		},
		nil,
	)

	requestBody := requests.RegisterRequest{
		Name:            "new user",
		Email:           "newUser@mail.com",
		Password:        "newpassword",
		ConfirmPassword: "newpassword",
		Currency:        "IDR",
	}

	reqJson, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	errHand := handler.Register(c)

	assert.NoError(t, errHand)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "oke", response["message"])
	assert.NotNil(t, response["data"])
}

func TestAuthHandler_Register_invalidJson(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Service should NOT be called
	service.AssertNotCalled(t, "Login")
}

func TestAuthHandler_Register_ValidationError(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)

	reqBody := requests.RegisterRequest{
		Name:            "new user",
		Email:           "invalidemail.com",
		Password:        "password",
		ConfirmPassword: "password",
		Currency:        "IDR",
	}

	request, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	c := e.NewContext(req, rec)

	err := handler.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}

	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Nil(t, response["data"])
	assert.Contains(t, response["message"], "email")
}

func TestAuthHandler_Register_RegisterFailed(t *testing.T) {
	service := new(MockAuthService)
	handler := NewAuthHandler(service)
	ctx := context.Background()
	service.On("Register", mock.AnythingOfType("*requests.RegisterRequest"), ctx).Return(&responses.RegisterResponse{}, &schemas.ResponseApiError{
		Status:  schemas.ApiErrorBadRequest,
		Message: "Email already in use",
	})

	requestBody := requests.RegisterRequest{
		Name:            "new user",
		Email:           "existinguser@example.com",
		Password:        "password",
		ConfirmPassword: "password",
		Currency:        "IDR",
	}

	reqJson, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	c := e.NewContext(req, rec)

	err := handler.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}

	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Nil(t, response["data"])
	assert.Equal(t, "Email already in use", response["message"])
}
