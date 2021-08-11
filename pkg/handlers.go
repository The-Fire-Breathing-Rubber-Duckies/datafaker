package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/db"
)

func GetTableData(c echo.Context) error {
	params := db.ConnectParams{
		Host:     c.QueryParam("host"),
		Port:     c.QueryParam("port"),
		User:     c.QueryParam("user"),
		Password: c.QueryParam("password"),
		Dbname:   c.QueryParam("dbname"),
		Sslmode:  c.QueryParam("sslmode"),
	}
	var schema db.Schema
	schema.Tables = make(map[string]db.TableSchema)
	database := db.Connect(params)
	tables := db.GetTables(database)
	for _, table := range tables {
		schema.Tables[table] = db.DescribeTable(database, table)
	}
	return c.JSON(http.StatusOK, schema)
}
