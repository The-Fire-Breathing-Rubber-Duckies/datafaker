package pkg

import (
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/describe", GetTableData)
	return e
}
