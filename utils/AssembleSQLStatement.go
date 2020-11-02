package utils

import "fmt"

func AssembleSQLStatement(tableName string, myMap map[int][]string) string{
	stmtStr := "INSERT INTO " + tableName + " VALUES("
	for range myMap[0] {
		fmt.Println("ok")
		stmtStr += "?,"
	}
	stmtStr = stmtStr[0:len(stmtStr)-1]
	stmtStr += ")"
	fmt.Println(stmtStr)
	return stmtStr
}
