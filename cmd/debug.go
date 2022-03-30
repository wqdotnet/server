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
	"log"
	"server/gserver/commonstruct"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "控制台",
	Long:  `gen sever ping GameServer `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			debug(args[0], args[1])
		} else {
			debug(strconv.Itoa(int(commonstruct.ServerCfg.ServerID)), "127.0.0.1")
		}
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}

func debug(serverid, ip string) {
	ok, servername := ping(serverid, ip)
	if !ok {
		return
	}

	for {
		command := strings.TrimSpace(prompt.Input("["+servername+"] > ", completer))
		if command == "quit" {
			return
		}
		if command == "" {
			break
		}
		term, err := call(strings.Split(command, " ")...)
		log.Printf("info: %v [%v]\n", term, command)

		if command == "shutdown" && err == nil {
			return
		}
	}

}

//commands   = []string{"ping", "reloadCfg", "state", "shutdown"}
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "quit", Description: "退出连接模式"},
		{Text: "state", Description: "查看服务器状态"},
		{Text: "reloadCfg", Description: "重新加载配置文件"},
		{Text: "shutdown", Description: "关闭服务器!!!"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
