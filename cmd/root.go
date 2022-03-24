package cmd

import (
	"fmt"
	"os"
	"server/gserver/commonstruct"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//"github.com/joho/godotenv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "root demo",
	Short: "root Short",
	Long:  `服务器`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var cfgfile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//initConfig()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgfile, "config", "", "config file (default is $HOME/cfg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgfile != "" {
		fmt.Println("initConfig config :", cfgfile)
		// Use config file from the flag.
		viper.SetConfigFile(cfgfile)
	} else {
		// Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// fmt.Println("initConfig config home:", home)
		// // Search config in home directory with name ".demo" (without extension).
		// viper.AddConfigPath(home)

		// dir, _ := os.Getwd()
		// viper.AddConfigPath(dir)
		// fmt.Println("initConfig config dir:", dir)
		viper.AddConfigPath(".")
		viper.SetConfigName("cfg")
	}

	//viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&commonstruct.ServerCfg)
		logrus.Info("Using config file:", viper.ConfigFileUsed())
	}
}
