package cmd

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cronConfigCmd represents the cronConfig command
var cronConfigCmd = &cobra.Command{
	Use:   "cronConfig",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cronConfig()
	},
}

func init() {
	rootCmd.AddCommand(cronConfigCmd)

}

func cronConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("cron")

	if err2 := viper.ReadInConfig(); err2 == nil {
		cron := cronCfg{}
		viper.Unmarshal(&cron)
		for k, v := range cron.Cronlist {
			fmt.Println(k, ":", v)
		}
	}
	//tool.ExamplePerlin_Noise2D()
}

func crontest() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	select {}
}

type cronCfg struct {
	Cronlist map[string]string
	// cronConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
