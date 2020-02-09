package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"bitbucket.com/libertywireless/icecream_manager/controller"
	"bitbucket.com/libertywireless/icecream_manager/lib"
	"net/http"
)

func loadRoutes(router *echo.Echo) {
	middleware.ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "you don't have permission to to this")

	router.GET("/", controller.AppHandler)
	v1 := router.Group("/api/v1")
	v1.GET("/products", controller.ProductListHandler)
	v1.GET("/products/:id", controller.ProductGetByIDHandler)
	auth := v1.Group("", middleware.JWT([]byte(lib.Config.Secret)))
	auth.POST("/products", controller.ProductAddHandler)
	auth.PUT("/products/:id", controller.ProductUpdateHandler)
	auth.DELETE("/products/:id", controller.ProductDeleteHandler)
}
