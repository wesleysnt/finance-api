package authhandler

import "github.com/labstack/echo/v4"

func Route(route *echo.Group) {
	handler := NewAuthHandler()
	auth := route.Group("/auth")

	auth.POST("/login", handler.Login)
}
