package testhandler

import "github.com/labstack/echo/v4"

func Route(route *echo.Group) {
	handler := NewTestHandler()
	test := route.Group("/test")

	test.GET("", handler.Test)
}
