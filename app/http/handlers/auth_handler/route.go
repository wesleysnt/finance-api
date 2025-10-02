package authhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/repositories"
	"github.com/wesleysnt/finance-api/app/services"
)

func Route(route *echo.Group) {
	userRepository := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepository)
	handler := NewAuthHandler(authService)
	auth := route.Group("/auth")

	auth.POST("/login", handler.Login)
}
