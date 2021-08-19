package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Insert(db *sql.DB, table string, data map[string]interface{}) (res sql.Result, ok bool) {
	// n := len(data)
	i := 0

	columnQ := ""
	valueQ := ""
	var columns []interface{}
	var values []interface{}
	for k, v := range data {
		if i > 0 {
			columnQ += ", "
			valueQ += ", "
		}
		// columnQ += "$" + strconv.Itoa(i+1) + ""
		// valueQ += "$" + strconv.Itoa(i+n+1) + ""
		columnQ += k
		valueQ += `'` + v.(string) + `'`
		columns = append(columns, k)
		values = append(values, v)
		i += 1
	}

	stmt := `
		INSERT INTO ` + table + `
			(` + columnQ + `)
		VALUES
			(` + valueQ + `)
	`

	// q, err := db.Prepare(stmt)

	// if err != nil {
	// 	log.Printf("%v", err)
	// 	log.Fatalf("Could not pre pare statement, %s", stmt)
	// }
	// defer q.Close()

	// input := []interface{}{}
	// input = append(append(input, columns...), values...)
	res, err := db.Exec(stmt)
	if err != nil {
		fmt.Printf("%s", stmt)
		log.Fatal(err)
	}

	fmt.Printf("%v", res)
	return res, true
}
