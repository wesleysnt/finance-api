package routes

import (
	"github.com/labstack/echo/v4"
	authhandler "github.com/wesleysnt/finance-api/app/http/handlers/auth_handler"
	testhandler "github.com/wesleysnt/finance-api/app/http/handlers/test_handler"
	"github.com/wesleysnt/finance-api/pkg"
)

func RegisterRoute(e *echo.Echo) {
	api := e.Group("api")
	v1 := api.Group("/v1")

	testhandler.Route(v1)
	authhandler.Route(v1)

	pkg.ListRouters(e)

}
