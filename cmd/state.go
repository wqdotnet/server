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
	"server/gserver/commonstruct"
	"strconv"

	"github.com/spf13/cobra"
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "获取服务器运行状态",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var serverid, ip string
		if len(args) == 2 {
			serverid = args[0]
			ip = args[1]
		} else {
			serverid = strconv.Itoa(int(commonstruct.ServerCfg.ServerID))
			ip = "127.0.0.1"
		}

		startDebugGen(serverid, ip)
		if info, err := call("state"); err == nil {
			fmt.Printf(" %v \n", info)
		} else {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(stateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
