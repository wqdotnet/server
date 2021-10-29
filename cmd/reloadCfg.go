/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"server/gserver"
	"strconv"

	"github.com/spf13/cobra"
)

// upcfgCmd represents the upcfg command
var upcfgCmd = &cobra.Command{
	Use:   "reloadcfg",
	Short: "重新加载配置",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var pong bool

		if len(args) == 2 {
			pong = ping(args[0], args[1])
		} else {
			pong = ping(strconv.Itoa(int(gserver.ServerCfg.ServerID)), "127.0.0.1")
		}

		if pong {
			if info, err := call("ReloadCfg"); err == nil {
				fmt.Printf("[%v] ReloadCfg  \n", info)
			} else {
				fmt.Print("err:", err)
			}
		} else {
			fmt.Printf(" not running \n")
		}
	},
}

func init() {
	rootCmd.AddCommand(upcfgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upcfgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upcfgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
