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
	"server/gserver/cfg"

	"github.com/spf13/cobra"
)

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		excel()
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func excel() {
	cfg.InitViperConfig("config", "json")
	// cfg.CheckBigMapConfig()

	// f := excelize.NewFile()
	// // Create a new sheet.
	// index := f.NewSheet("AreasList")

	// f.DeleteSheet("Sheet1")

	// // Set value of a cell.
	// //f.SetCellValue("AreasList", "B2", 100)
	// f.SetCellValue("AreasList", "A1", 100)

	// areaslist := cfg.GameCfg.MapInfo.Areas

	// f.SetCellValue("AreasList", "A1", "AreasIndex")
	// f.SetCellValue("AreasList", "A2", "int")
	// f.SetCellValue("AreasList", "A3", "区域城池索引")
	// f.SetCellValue("AreasList", "B1", "Beside")
	// f.SetCellValue("AreasList", "B2", "int")
	// f.SetCellValue("AreasList", "B3", "相邻城池")

	// for num, arecfg := range areaslist {
	// 	// index := cfg.GameCfg.MapInfo.IndexCfg[arecfg.Setindex-1]
	// 	// tmpareas := GameCfg.MapInfo.Areas[index-1]
	// 	f.SetCellValue("AreasList", fmt.Sprintf("A%v", num+4), arecfg.Setindex)

	// 	beside := []int{}
	// 	for _, v := range arecfg.Beside {
	// 		tmpbaside := cfg.GameCfg.MapInfo.Areas[v-1]
	// 		beside = append(beside, tmpbaside.Setindex)
	// 	}
	// 	f.SetCellValue("AreasList", fmt.Sprintf("B%v", num+4), beside)

	// }

	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("bigmap.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

}
