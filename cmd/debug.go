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
	"os"
	"path/filepath"
	"server/gserver/commonstruct"
	"strconv"
	"strings"

	"github.com/peterh/liner"
	"github.com/spf13/cobra"
)

var (
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
	names      = []string{"ping", "reloadCfg", "state", "shutdown"}
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

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(history_fn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		if name, err := line.Prompt("[" + servername + "] -> "); err == nil {
			term, err := call(strings.Split(name, " ")...)

			if err != nil {
				log.Print("err: ", name)
			} else {
				log.Printf("info: %v \n", term)
				line.AppendHistory(name)
			}

			if name == "shutdown" && err == nil {
				return
			}

		} else if err == liner.ErrPromptAborted {
			if f, err := os.Create(history_fn); err != nil {
				log.Print("Error writing history file: ", err)
			} else {
				line.WriteHistory(f)
				f.Close()
			}
			log.Print("Aborted")
			return
		} else {
			log.Print("Error reading line: ", err)
		}
	}

}
