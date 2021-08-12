package db

type TableMetaData struct {
	Name            string
	ColumnsMetaData map[string]ColumnMetaData
}

type ColumnMetaData struct {
	DataDescription string `example: "phone, email, address, etc..."`
	FieldToRepeat   string
	RepeatAmount    string
}
