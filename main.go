package main

import (
	"github.com/gookit/gcli/v3"
	"pure-admin-cli/cmd"
)

var (
	version = "unknown"
)

func main() {
	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = version
		app.Desc = "create pure-admin cli application"
		app.Name = "pure"
	})
	app.AddCommand(cmd.InitCmd)
	app.AddCommand(cmd.NewCmd)
	app.Run(nil)
}
