package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func ExecuteInsert(myMap map[int][]string, stmtStr string, db *sql.DB, fieldInfo []Field) {
	fmt.Println("-----------------START INSERT-----------------")
	//开启事务，因为要一次性执行很多insert语句，所以在全部读入后再commit。否则，每执行一条insert语句就要开启一次事务，速度会很慢！
	conn, err := db.Begin()
	if err != nil {
		return
	}
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	//遍历从csv/sqlite文件中读入的数据
	for _, value := range myMap {
		colPtrs := make([]interface{}, len(value))
		for i, val := range value{
			//根据目标表结构对数据进行类型转换
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
	//提交事务
	if err := conn.Commit(); err != nil {
		panic(err)
	}
	fmt.Println("-----------------END INSERT-----------------")
}
