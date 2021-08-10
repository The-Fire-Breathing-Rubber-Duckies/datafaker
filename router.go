package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>Hello World</h1>`)
	})
	return e
}
