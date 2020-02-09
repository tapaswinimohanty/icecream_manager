package app

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
	"bitbucket.com/libertywireless/icecream_manager/database"
	"bitbucket.com/libertywireless/icecream_manager/lib"
	"time"
)

// Start app
func Run() {
	lib.LoadConfig("config.yaml")
	lib.ConnectDatabase()
	database.Migrate()

	router := echo.New()
	router.Use(middleware.Gzip())
	router.Use(middleware.Recover())
	router.Use(middleware.RemoveTrailingSlash())
	router.Static("/", "public")

	router.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &lib.CustomContext{Context: c}
			return h(cc)
		}
	})

	// load all routes
	loadRoutes(router)


	router.Server.Addr = fmt.Sprintf(`:%v`, lib.Config.AppPort)
	router.Logger.Fatal(graceful.ListenAndServe(router.Server, 5*time.Second))
}
