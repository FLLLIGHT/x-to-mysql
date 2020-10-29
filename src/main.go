package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func main() {
	tableName := "room"
	db, err := sql.Open("mysql", "root:541978@/test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

	fieldInfo := ParseMySQLTableSchema("test", "room", db)

	fmt.Println(len(fieldInfo))
	for _, item := range fieldInfo {
		fmt.Println(item)
	}

	myMap := ReadFromSQLite("xxxdatabase.db", "room")

	stmtStr := "INSERT INTO " + tableName + " VALUES("
	for _, _ = range myMap[0] {
		fmt.Println("ok")
		stmtStr += "?,"
	}
	stmtStr = stmtStr[0:len(stmtStr)-1]
	stmtStr += ")"
	fmt.Println(stmtStr)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		panic(err)
	}

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
				colPtrs[i] = val
			}
		}
		_, err = stmt.Exec(colPtrs...)
		if err != nil {
			panic(err)
		}
	}

}