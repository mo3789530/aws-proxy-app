package cmd

import (
	"aws-proxy-app/internal/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// serverCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin},
		}))
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
			Skipper: func(c echo.Context) bool {
				if c.Request().URL.Path == "/" {
					return true
				}
				return false
			},
		}))
		e.Use(middleware.Recover())
		//cnf := proxy.NewConfig(os.Getenv("CONFIG_PATH"))
		server.SetupRootRoutes(e)
		server.SetupProxyRoutes(e)

		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
