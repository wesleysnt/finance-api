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

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	requests := &requests.LoginRequest{}

	if err := c.Bind(requests); err != nil {
		return helpers.ResponseApiError(c, err.Error(), 400, nil)
	}

	if err := c.Validate(requests); err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	resp, err := h.service.Login(requests, c.Request().Context())

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "oke", resp)
}

func (h *AuthHandler) Register(c echo.Context) error {
	requests := &requests.RegisterRequest{}

	if err := c.Bind(requests); err != nil {
		return helpers.ResponseApiError(c, err.Error(), 400, nil)
	}

	if err := c.Validate(requests); err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	resp, err := h.service.Register(requests, c.Request().Context())

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "oke", resp)
}
