package authhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/repositories"
	"github.com/wesleysnt/finance-api/app/services"
	"github.com/wesleysnt/finance-api/pkg"
	"github.com/wesleysnt/finance-api/pkg/auth"
)

func Route(route *echo.Group) {
	db := pkg.Orm()
	jwt := auth.NewJWTService()
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository, jwt)
	handler := NewAuthHandler(authService)
	auth := route.Group("/auth")

	auth.POST("/login", handler.Login)
	auth.POST("/register", handler.Register)
}
