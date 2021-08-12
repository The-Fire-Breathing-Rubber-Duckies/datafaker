package db

import (
	"database/sql"
	"fmt"
	"sort"
)

func WriteToDb(db *sql.DB, schema TableSchema, table string) error {
	var err error
	sort.Slice(schema.Columns[:], func(i, j int) bool {
		return schema.Columns[i].OrdinalPos < schema.Columns[j].OrdinalPos
	})
	query := fmt.Sprintf(`INSERT INTO %s (`, table)
	for i, column := range schema.Columns {
		if i == len(schema.Columns)-1 {
			query += column.ColumnName
		} else {
			query += fmt.Sprintf(`%s, `, column.ColumnName)
		}
	}
	query += ")\n VALUES("
	query += `'dfc43e30-0ac6-4528-80e9-cad3a5ca01ca', '1c867ca2-3b92-45d4-bbb0-f95b4ad5d696', 1118, 'no_auth', 'nesciunt', '584084f9-e305-4d9c-8303-4ee4ce89f7be', 'da0fd70d-57bc-4cbe-8dcf-31f67f544a60', 'c1aa444a-f5f6-4141-89e1-0ebe8991f731', 'update', 'decrypt', '2009-05-24 10:21:44', '{bVU,Tkq,sNe}', '{LsScq,ilvXd,fwJcp,QOhZt,nkhBf}', '{wscnOLf,xIMpeKs,XiQuNBC,VOvWXgg,kDImbJy,fGYufXX,LdNxceu}')`
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	return err
}
