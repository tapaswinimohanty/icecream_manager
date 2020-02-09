package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"bitbucket.com/libertywireless/icecream_manager/lib"
)

func AppHandler(c echo.Context) error {
	return c.String(200, fmt.Sprintf(`%v v%v is working on /api`,
		lib.Config.AppName,
		lib.Config.AppVersion))
}
