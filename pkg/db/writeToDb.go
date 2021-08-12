package db

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/brianvoe/gofakeit/v6"
)

func WriteToDb(db *sql.DB, schema TableSchema, metaData TableMetaData) error {
	gofakeit.Seed(0)
	var err error
	sort.Slice(schema.Columns[:], func(i, j int) bool {
		return schema.Columns[i].OrdinalPos < schema.Columns[j].OrdinalPos
	})
	query := fmt.Sprintf(`INSERT INTO %s (`, metaData.Name)
	for i, column := range schema.Columns {
		if i == len(schema.Columns)-1 {
			query += column.ColumnName
		} else {
			query += fmt.Sprintf(`%s, `, column.ColumnName)
		}
	}
	query += ")\n VALUES("
	for i, column := range schema.Columns {
		// TODO: look into mapping the gofakeit functions in a map inside of the metadata model and then forcing the user into choosing metadata values that match the function names so that all we have to do is call the function without doing any checking.
		fmt.Println(i, column)
	}
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	return err
}
