package lib

import (
	"github.com/labstack/echo"
	"net/http"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) OK(v interface{}) error {
	return c.JSON(http.StatusOK, &v)
}

func (c *CustomContext) BadRequest(message string) error {
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": message,
	})
}

func (c *CustomContext) NotFound(message string) error {
	return c.JSON(http.StatusNotFound, echo.Map{
		"message": message,
	})
}

func (c *CustomContext) Conflict(message string) error {
	return c.JSON(http.StatusConflict, echo.Map{
		"message": message,
	})
}

func (c *CustomContext) Internal(message string) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"message": message,
	})
}
