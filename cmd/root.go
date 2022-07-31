package cmd

import "github.com/gookit/gcli/v3"

var (
	version = "unknown"
	banner  = `┏━┃┃ ┃┏━┃┏━┛  ┏━┃┏━ ┏┏ ┛┏━   ┏━┛┃  ┛
┏━┛┃ ┃┏┏┛┏━┛  ┏━┃┃ ┃┃┃┃┃┃ ┃  ┃  ┃  ┃
┛  ━━┛┛ ┛━━┛  ┛ ┛━━ ┛┛┛┛┛ ┛  ━━┛━━┛┛
`
)

func Execute() error {
	app := gcli.New()
	app.Name = "pure-cli"
	app.Version = version
	app.Desc = "pure admin cli"
	app.Logo.Text = banner
	app.AddCommand(initCmd)
	app.AddCommand(newCmd)
	app.Run(nil)
	return nil
}
