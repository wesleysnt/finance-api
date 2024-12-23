package routes

import (
	"github.com/labstack/echo/v4"
	testhandler "github.com/wesleysnt/go-base/app/http/handlers/test_handler"
	"github.com/wesleysnt/go-base/pkg"
)

func RegisterRoute(e *echo.Echo) {
	api := e.Group("api")
	v1 := api.Group("/v1")

	testhandler.Route(v1)

	pkg.ListRouters(e)

}
