package testhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/helpers"
	"github.com/wesleysnt/finance-api/app/schemas"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c echo.Context) error {
	er := c.QueryParam("error")
	var err error

	if er == "true" {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "error",
		}
	}

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiOk(c, "oke", nil)
}
