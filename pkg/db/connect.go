package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectParams struct {
	Host     string `localhost`
	Port     int    `default:"5432"`
	User     string
	Password string
	Dbname   string
	Sslmode  string `default:"disable"`
}

func Connect(p ConnectParams) (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Dbname, p.Sslmode)

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return DB
}
