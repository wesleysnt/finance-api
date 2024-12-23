package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/wesleysnt/go-base/app/config"
	"github.com/wesleysnt/go-base/app/routes"
	"github.com/wesleysnt/go-base/cmd/commands"
)

func main() {
	env := config.GetEnv()
	config.ConnectDB(env.Database)

	if len(os.Args) >= 2 {
		commands.Execute()
		return
	}

	e := echo.New()

	routes.RegisterRoute(e)
	e.HideBanner = true
	e.Logger.Fatal(e.Start(env.Server.Host + ":" + env.Server.Port))

}
