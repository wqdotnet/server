package cmd

import (
	"server/db"
	"server/gserver"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清理数据",
	Long:  `清理 redis 缓存数据  [areasSMap] [troopsSMap]`,
	Run: func(cmd *cobra.Command, args []string) {
		clean()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func clean() {
	//db.StartMongodb(gserver.ServerCfg.Mongodb, gserver.ServerCfg.MongoConnStr)
	//client, database := getDatabase()

	db.StartRedis(gserver.ServerCfg.RedisConnStr, gserver.ServerCfg.RedisDB)
	db.RedisExec("del", "areasSMap")
	db.RedisExec("del", "troopsSMap")
}
