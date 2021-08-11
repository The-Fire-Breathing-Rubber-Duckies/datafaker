package db

import (
	"database/sql"
	"log"
)

func GetTables(db *sql.DB) (tables []string) {
	q, err := db.Prepare(`
			SELECT
				tablename
			FROM pg_catalog.pg_tables
			WHERE schemaname != 'pg_catalog' AND 
					schemaname != 'information_schema';
	 `)

	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	rows, err := q.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string

		if err := rows.Scan(
			&table,
		); err != nil {
			log.Fatal(err)
		}

		tables = append(tables, table)
	}

	return tables
}
