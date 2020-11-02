package utils

func AssembleSQLStatement(tableName string, length int) string{
	stmtStr := "INSERT INTO " + tableName + " VALUES("
	for i:=0; i<length; i++ {
		stmtStr += "?,"
	}
	stmtStr = stmtStr[0:len(stmtStr)-1]
	stmtStr += ")"
	return stmtStr
}
