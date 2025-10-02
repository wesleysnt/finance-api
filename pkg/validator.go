package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/wesleysnt/finance-api/app/schemas"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}
