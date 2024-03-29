package cmd

import (
	"context"
	"server/db"
	"server/gserver/commonstruct"

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
	db.StartMongodb(commonstruct.ServerCfg.Mongodb, commonstruct.ServerCfg.MongoConnStr)
	_, database := db.GetDatabase()
	database.Drop(context.Background())

	db.StartRedis(commonstruct.ServerCfg.RedisConnStr, commonstruct.ServerCfg.RedisDB)
	//db.RedisExec("del", "ConnectNumber")
	db.RedisExec("FLUSHDB", "")
}
