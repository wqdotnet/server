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
	"strconv"
	"strings"

	"github.com/go-basic/uuid"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "excel [数量] [长度] 生成字符串cdk",
	Long:  `生成excel 字符串cdk`,
	Args:  cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		num, _ := strconv.Atoi(args[0])
		lennum, _ := strconv.Atoi(args[1])
		excel(num, lennum)
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)
}

func excel(num, lennum int) {
	fmt.Println("生成excel")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("CDKList")

	f.DeleteSheet("Sheet1")

	tmpmap := map[string]interface{}{}

	i := 0
	for {
		uuid := uuid.New()
		strNum := strings.Replace(uuid, "-", "", -1)[0:lennum]

		if _, ok := tmpmap[strNum]; ok {
			continue
		}

		i++
		f.SetCellValue("CDKList", fmt.Sprintf("A%v", i), strNum)
		f.SetCellValue("CDKList", fmt.Sprintf("B%v", i), uuid)
		if i >= num {
			break
		}
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(fmt.Sprintf("%v个随机长度为%v的cdk.xlsx", num, lennum)); err != nil {
		fmt.Println(err)
	}

}
