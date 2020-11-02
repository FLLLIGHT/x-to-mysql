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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	)

var (
	username string
	password string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "x-to-mysql",
	Short: "x-to-mysql is a useful tool for you to export data from csv/sqlite to mysql",
	Long: `x-to-mysql is a useful tool for you to export data from csv/sqlite to mysql`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username of your mysql database")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of your mysql database")

	err := rootCmd.MarkPersistentFlagRequired("username")
	if err != nil {
		panic(err)
	}
	err = rootCmd.MarkPersistentFlagRequired("password")
	if err != nil {
		panic(err)
	}
}
