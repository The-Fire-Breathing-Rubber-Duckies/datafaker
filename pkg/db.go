package pkg

import (
	"database/sql"
	"fmt"
)

type Result struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

type ConnectParams struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Sslmode  string `default:"disable"`
}

func Connect(p ConnectParams) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Dbname, p.Sslmode)

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	rows, err := DB.Query(`SELECT c.column_name, c.data_type, c.is_nullable
	FROM information_schema.table_constraints tc 
	JOIN information_schema.constraint_column_usage AS ccu USING (constraint_schema, constraint_name) 
	JOIN information_schema.columns AS c ON c.table_schema = tc.constraint_schema
		AND tc.table_name = c.table_name AND ccu.column_name = c.column_name
	WHERE constraint_type = 'PRIMARY KEY' and tc.table_name = 'tdf_transactions'`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var result Result
		err = rows.Scan(&result.Field, &result.Type, &result.Null)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

}
