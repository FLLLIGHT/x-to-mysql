/*
Copyright © 2020 Xuan Zitao <18302010015@fudan.edu.cn>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"

	"github.com/FLLLIGHT/x-to-mysql/utils"
	"github.com/spf13/cobra"
)

var (
	dataSource string
	fromTableName string
	toTableName string
	toDatabase string
	encoding string
)

// sqliteCmd represents the sqlite command
var sqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "将数据从sqlite数据库导入至mysql",
	Long: `使用sqlite指令，可以从指定的sqlite数据库(即db文件)中读取指定的表中的数据，并将数据插入指定的MySQL数据库的table中`,
	Run: func(cmd *cobra.Command, args []string) {
		//连数据库
		db := utils.ConnectToMySQL(username, password, toDatabase)
		defer db.Close()
		//读目标表的表结构
		fieldInfo := utils.ParseMySQLTableSchema(toDatabase, toTableName, db)
		//读sqlite数据库
		myMap := ReadFromSQLite(dataSource, fromTableName)
		//根据目标表的表结构组装prepare statement
		stmtStr := utils.AssembleSQLStatement(toTableName, len(fieldInfo))
		//执行插入语句
		utils.ExecuteInsert(myMap, stmtStr, db, fieldInfo)
	},
}

func init() {
	rootCmd.AddCommand(sqliteCmd)

	sqliteCmd.Flags().StringVarP(&dataSource, "source", "s", "", "name of your db file (with extension)")
	sqliteCmd.Flags().StringVarP(&toTableName, "toTable", "t", "", "name of your mysql table")
	sqliteCmd.Flags().StringVarP(&fromTableName, "fromTableName", "f", "", "name of your sqlite table")
	sqliteCmd.Flags().StringVarP(&toDatabase, "toDatabase", "d", "", "name of your mysql database")

	err := sqliteCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = sqliteCmd.MarkFlagRequired("toTable")
	if err != nil {
		panic(err)
	}
	err = sqliteCmd.MarkFlagRequired("fromTableName")
	if err != nil {
		panic(err)
	}
	err = sqliteCmd.MarkFlagRequired("toDatabase")
	if err != nil {
		panic(err)
	}
}

func ReadFromSQLite(dataSourceName, tableName string) map[int][]string{
	fmt.Println("-----------------START READ FROM SQLite-----------------")

	db, err := sql.Open("sqlite3", dataSourceName)
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

