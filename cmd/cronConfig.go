package cmd

import (
	"fmt"
	"server/tool"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cronConfigCmd represents the cronConfig command
var cronConfigCmd = &cobra.Command{
	Use:   "cron",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cronConfig()
		l := make([]chan int32, 100000)
		for i := 0; i < 10000; i++ {
			cs := make(chan int32)
			l[i] = cs
			go func() {
				for i := range cs {
					if i == 1 || i == 100 || i == 9999 || i == 99999 {
						fmt.Println(tool.GoID(), ":", i)
					}
				}
			}()

		}
		fmt.Println("start cron ")
		crontest(l)
		select {}
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

func crontest(sendchan []chan int32) {

	c := cron.New(cron.WithSeconds())

	// cron.New(cron.WithChain(
	// 	cron.Recover(logger),  // or use cron.DefaultLogger
	//   ))
	i := int32(0)

	spec := "* * * * * ?"
	c.AddFunc(spec, func() {
		i++
		for sc, c := range sendchan {
			c <- int32(sc)
		}
	})
	c.Start()

	select {}
}

type cronCfg struct {
	Cronlist map[string]string
	// cronConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
