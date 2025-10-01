package authhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/helpers"
	"github.com/wesleysnt/finance-api/app/http/requests"
	"github.com/wesleysnt/finance-api/app/schemas"
	"github.com/wesleysnt/finance-api/app/services"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: *services.NewAuthService(),
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	requests := &requests.LoginRequest{}

	if err := c.Bind(requests); err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}

	resp, err := h.service.Login(requests)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "oke", resp)
}
