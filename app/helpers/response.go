package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/schemas"
)

func CatchErrorResponseApi(r *schemas.ResponseApiError) *schemas.SetResponseApiError {
	switch r.Status {
	case schemas.ApiErrorBadRequest:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusBadRequest,
			Message:    r.Message,
		}
	case schemas.ApiErrorForbidden:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusForbidden,
			Message:    r.Message,
		}
	case schemas.ApiErrorNotFound:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusNotFound,
			Message:    r.Message,
		}
	case schemas.ApiErrorUnprocessAble:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    r.Message,
		}
	default:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    r.Message,
		}
	}
}

func ResponseApi(ctx echo.Context, msg string, status string, statusCode int, data any, errors any) error {
	details := schemas.DetailResponse{
		StatusCode: statusCode,
		Path:       ctx.Request().RequestURI,
		Method:     string(ctx.Request().Method),
		Status:     status,
	}

	if statusCode >= 400 {
		return ctx.JSON(statusCode, schemas.ResponseApi{
			Valid:   false,
			Message: msg,
			Data:    data,
			Errors:  errors,
			Details: details,
		})
	}

	return ctx.JSON(statusCode, schemas.ResponseApi{
		Valid:   true,
		Message: msg,
		Data:    data,
		Errors:  errors,
		Details: details,
	})
}

func ResponseApiCreated(ctx echo.Context, msg string, data any) error {
	return ResponseApi(ctx, msg, "success_created", http.StatusCreated, data, nil)
}

func ResponseApiOk(ctx echo.Context, msg string, data any) error {
	return ResponseApi(ctx, msg, "success_ok", http.StatusOK, data, nil)
}

func ResponseApiUnauthorized(ctx echo.Context, msg string) error {
	return ResponseApi(ctx, msg, "error_unauthorized", http.StatusUnauthorized, nil, nil)
}

func ResponseApiForbidden(ctx echo.Context, msg string) error {
	return ResponseApi(ctx, msg, "error_forbidden", http.StatusForbidden, nil, nil)
}

func ResponseApiBadRequest(ctx echo.Context, msg string, errors any) error {
	return ResponseApi(ctx, msg, "error_bad_request", http.StatusBadRequest, nil, errors)
}

func ResponseApiError(ctx echo.Context, msg string, statusCode int, errors any) error {
	return ResponseApi(ctx, msg, "error_api", statusCode, nil, errors)
}
