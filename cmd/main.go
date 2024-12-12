package main

import (
	"os"

	"github.com/wesleysnt/go-base/app/config"
	"github.com/wesleysnt/go-base/cmd/commands"
)

func main() {
	env := config.GetEnv()
	config.ConnectDB(env.Database)

	if len(os.Args) >= 2 {
		commands.Execute()
		return
	}

}
