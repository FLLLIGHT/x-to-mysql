package main

import (
	"database/sql"
	"fmt"
)

type Field struct {
	fieldName string
	dataType string
	isNullable string
	length int
}

func ParseMySQLTableSchema(dbName string, tableName string, db *sql.DB) [] Field{
	fmt.Println("-----------------START PARSE MySQL TABLE SCHEMA-----------------")
	stmt, err := db.Prepare("SELECT COLUMN_NAME, DATA_TYPE, " +
		"IS_NULLABLE, IFNULL(CHARACTER_MAXIMUM_LENGTH, 0) " +
		"FROM information_schema.columns " +
		"WHERE table_schema = ? AND table_name = ?")

	if err != nil {
		panic(err.Error())
	}

	var fields []Field

	rows, err := stmt.Query(dbName, tableName)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var f Field
		err = rows.Scan(&f.fieldName, &f.dataType, &f.isNullable, &f.length)
		if err != nil {
			panic(err.Error())
		}
		fields = append(fields, f)
	}
	fmt.Println("-----------------END PARSE MySQL TABLE SCHEMA-----------------")
	return fields
}