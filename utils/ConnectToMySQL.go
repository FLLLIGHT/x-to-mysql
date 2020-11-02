package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func ConnectToMySQL(username, password, tableName string) *sql.DB {
	db, err := sql.Open("mysql", username+":"+password+"@/"+tableName+"?multiStatements=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}