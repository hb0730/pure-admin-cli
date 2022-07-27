package cmd

import "github.com/gookit/gcli/v3"

var (
	version = "unknown"
	//rootCmd = &cobra.Command{
	//	Use:     "pure",
	//	Short:   "pure admin cli",
	//	Version: version,
	//}
)

func init() {
	//rootCmd.AddCommand(initCmd)
	//rootCmd.AddCommand(newCmd)
}

func Execute() error {
	//return rootCmd.Execute()
	app := gcli.New()
	app.Name = "pure-cli"
	app.Version = version
	app.Desc = "pure admin cli"
	app.AddCommand(initCmd)
	app.AddCommand(newCmd)
	app.Run(nil)
	return nil
}
