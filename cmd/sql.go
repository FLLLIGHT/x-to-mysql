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
	"github.com/FLLLIGHT/x-to-mysql/utils"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "执行sql文件中的sql语句",
	Long: `使用sql命令，可以从指定的sql文件中读取sql语句，并在指定的MySQL数据库中执行这些语句。`,
	Run: func(cmd *cobra.Command, args []string) {
		//读sql文件
		sqlFile, err := os.Open(dataSource)
		if err != nil {
			panic(err)
		}
		defer sqlFile.Close()
		sqlBytes, err := ioutil.ReadAll(sqlFile)
		sqlStr := string(sqlBytes[:])

		//连数据库并执行sql语句
		db := utils.ConnectToMySQL(username, password, toDatabase)
		defer db.Close()
		_, err = db.Exec(sqlStr)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)

	sqlCmd.Flags().StringVarP(&dataSource, "source", "s", "", "name of your sql file (with extension)")
	sqlCmd.Flags().StringVarP(&toDatabase, "toDatabase", "d", "", "name of your mysql database")

	err := sqlCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = sqlCmd.MarkFlagRequired("toDatabase")
	if err != nil {
		panic(err)
	}
}
