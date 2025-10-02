package main

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/finance-api/app/config"
	"github.com/wesleysnt/finance-api/app/helpers"
	"github.com/wesleysnt/finance-api/app/routes"
	"github.com/wesleysnt/finance-api/cmd/commands"
	"github.com/wesleysnt/finance-api/pkg"
)

func main() {
	env := config.GetEnv()
	config.ConnectDB(env.Database)

	if len(os.Args) >= 2 {
		commands.Execute()
		return
	}

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	routes.RegisterRoute(e)
	e.RouteNotFound("/*", missingRouteHandler)
	e.HideBanner = true
	e.Logger.Fatal(e.Start(env.Server.Host + ":" + env.Server.Port))
}

func missingRouteHandler(c echo.Context) error {
	return helpers.ResponseApiError(c, "Route not found", 404, nil)
}
