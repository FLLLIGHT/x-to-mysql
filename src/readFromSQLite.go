package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

func ReadFromSQLite(dataSourceName, tableName string) map[int][]string{
	fmt.Println("-----------------START READ FROM SQLite-----------------")

	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	var myMap = make(map[int][]string)
	rows, err := db.Query("SELECT * FROM `"+tableName+"`;")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	count := 0
	cols := make([]string, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}
	for rows.Next() {
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}
		for _, col := range cols {
			myMap[count] = append(myMap[count], col)
		}
		count++
	}
	fmt.Println("-----------------END READ FROM SQLite-----------------")
	return myMap
}
