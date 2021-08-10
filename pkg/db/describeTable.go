package db

import (
	"database/sql"
	"log"
)

type TableSchema struct {
	Columns []Column
	Indexes []Index
}

type Column struct {
	ColumnName    string
	OrdinalPos    int
	ColumnDefault string
	IsNullable    string
	DataType      string
}

type Index struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func DescribeTable(db *sql.DB, tableName string) (tableSchema TableSchema) {
	descTable(db, tableName, &tableSchema)
	descIndexes(db, tableName, &tableSchema)
	return tableSchema
}

func descTable(db *sql.DB, tableName string, tableSchema *TableSchema) {
	q, err := db.Prepare(`
		SELECT 
			column_name,
			ordinal_position,
			column_default,
			is_nullable,
			data_type
		FROM 
			information_schema.columns
		WHERE 
			table_name = $1;
	 `)

	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	rows, err := q.Query(tableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var column Column

		if err := rows.Scan(
			&column.ColumnName,
			&column.OrdinalPos,
			&column.ColumnDefault,
			&column.IsNullable,
			&column.DataType,
		); err != nil {
			log.Fatal(err)
		}

		tableSchema.Columns = append(tableSchema.Columns, column)
	}
}

func descIndexes(db *sql.DB, tableName string, tableSchema *TableSchema) {
	q, err := db.Prepare(`
		SELECT
			c.column_name,
			c.data_type,
			c.is_nullable
		FROM information_schema.table_constraints tc 
		JOIN information_schema.constraint_column_usage AS ccu USING (constraint_schema, constraint_name) 
		JOIN information_schema.columns AS c ON c.table_schema = tc.constraint_schema
			AND tc.table_name = c.table_name AND ccu.column_name = c.column_name
		WHERE constraint_type = 'PRIMARY KEY' and tc.table_name = $1
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	var index Index

	if err := q.QueryRow(tableName).Scan(
		&index.Field,
		&index.Type,
		&index.Null,
	); err != nil {
		log.Fatal(err)
	}

	tableSchema.Indexes = append(tableSchema.Indexes, index)
}
