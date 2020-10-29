package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func main() {
	tableName := "student"
	db, err := sql.Open("mysql", "root:541978@/test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

	fieldInfo := ParseMySQLTableSchema("test", tableName, db)

	fmt.Println(len(fieldInfo))
	for _, item := range fieldInfo {
		fmt.Println(item)
	}

	myMap := ReadFromCSV("./student.csv")
	//myMap := ReadFromSQLite("xxxdatabase.db", tableName)

	stmtStr := "INSERT INTO " + tableName + " VALUES("
	for _, _ = range myMap[0] {
		fmt.Println("ok")
		stmtStr += "?,"
	}
	stmtStr = stmtStr[0:len(stmtStr)-1]
	stmtStr += ")"
	fmt.Println(stmtStr)
	conn, err := db.Begin()
	if err != nil {
		return
	}
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
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