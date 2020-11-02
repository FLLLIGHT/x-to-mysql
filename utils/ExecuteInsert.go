package utils

import (
	"database/sql"
	"strconv"
	"time"
)

func ExecuteInsert(myMap map[int][]string, stmtStr string, db *sql.DB, fieldInfo []Field) {
	conn, err := db.Begin()
	if err != nil {
		return
	}
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	//1815 47
	for _, value := range myMap {
		colPtrs := make([]interface{}, len(value))
		for i, val := range value{
			if fieldInfo[i].dataType == "int" {
				v, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				colPtrs[i] = v
			}
			if fieldInfo[i].dataType == "varchar" {
				colPtrs[i] = val
			}
			if fieldInfo[i].dataType == "datetime" {
				layout := "2006-01-02 15:04"
				colPtrs[i], err = time.ParseInLocation(layout, val, time.Local)
				if err != nil {
					panic(err)
				}
			}
		}
		if _, err := stmt.Exec(colPtrs...); err != nil{
			panic(err)
		}
	}
	if err := conn.Commit(); err != nil {
		panic(err)
	}
}
