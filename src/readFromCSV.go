package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadFromCSV(tableName string) map[int][]string {
	fmt.Println("-----------------START READ FROM CSV-----------------")
	var myMap = make(map[int][]string)
	csvFile, err := os.Open(tableName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	count := -1
	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if count == -1{
			for _, r := range record {
				fmt.Println(r)
			}
		} else {
			for _, col := range record {
				//cc, _ := iconv.ConvertString(col, "gb2312", "utf-8")
				myMap[count] = append(myMap[count], col)
			}
		}
		count++
	}
	return myMap
}
