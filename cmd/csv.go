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
	"encoding/csv"
	"fmt"
	"github.com/djimenez/iconv-go"
	"github.com/saintfish/chardet"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"x-to-mysql/utils"
)

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "将数据从csv文件导入至mysql",
	Long: `使用csv指令，可以从指定的csv文件中读取数据，并将数据插入指定的MySQL数据库的table中`,

	Run: func(cmd *cobra.Command, args []string) {
		//连数据库
		db := utils.ConnectToMySQL(username, password, toDatabase)
		defer db.Close()
		//读目标表的表结构
		fieldInfo := utils.ParseMySQLTableSchema(toDatabase, toTableName, db)
		//读csv文件
		myMap := ReadFromCSV(dataSource)
		//根据目标表的表结构组装prepare statement
		stmtStr := utils.AssembleSQLStatement(toTableName, len(fieldInfo))
		//执行插入语句
		utils.ExecuteInsert(myMap, stmtStr, db, fieldInfo)
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.Flags().StringVarP(&dataSource, "source", "s", "", "name of your csv file (with extension)")
	csvCmd.Flags().StringVarP(&toTableName, "toTable", "t", "", "name of your mysql table")
	csvCmd.Flags().StringVarP(&toDatabase, "toDatabase", "d", "", "name of your mysql database")
	csvCmd.Flags().StringVarP(&encoding, "encoding", "e", "", "charset encoding of your csv file (optional)")

	err := csvCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = csvCmd.MarkFlagRequired("toTable")
	if err != nil {
		panic(err)
	}
	err = csvCmd.MarkFlagRequired("toDatabase")
	if err != nil {
		panic(err)
	}
}

func ReadFromCSV(dataSource string) map[int][]string {
	fmt.Println("-----------------START READ FROM CSV-----------------")
	var myMap = make(map[int][]string)
	csvFile, err := os.Open(dataSource)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	var fromEncoding string
	//如果没有设置flag并指定字符集，则猜测字符集charset
	if encoding == "" {
		detector := chardet.NewTextDetector()
		all, _ := ioutil.ReadAll(csvFile)
		//将reader重新seek回头部，不然待会读不出了
		_, err = csvFile.Seek(0, 0)
		if err != nil {
			panic(err)
		}
		rs, _ := detector.DetectBest(all)

		if strings.HasPrefix(rs.Charset, "GB") {
			fromEncoding = "gbk"
		} else {
			fromEncoding = rs.Charset
		}
	} else {
		fromEncoding = encoding
	}
	fmt.Println("CSV Charset: "+fromEncoding)

	count := -1
	for{
		record, err := r.Read()
		//读到文件结尾为止
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// 跳过第一行
		if count != -1 {
			//遍历行内所有数据，append到对应位置
			for _, col := range record {
				if fromEncoding != "UTF-8" {
					col, _ = iconv.ConvertString(col, fromEncoding, "utf-8")
				}
				myMap[count] = append(myMap[count], col)
			}
		}
		//读下一行
		count++
	}
	return myMap
}