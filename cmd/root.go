package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// ServerConfig  server cfg
type ServerConfig struct {
	CfgFile string
	Host    string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{
	CfgFile: "",
	Host:    "localhost:8080",
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "Short",
	Long:  `long`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

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

	rootCmd.PersistentFlags().StringVar(&ServerCfg.CfgFile, "config", "", "config file (default is $HOME/demo.yaml)")
	rootCmd.PersistentFlags().StringVar(&ServerCfg.Host, "host", "", "localhost:8080")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("cfgFile:", ServerCfg.CfgFile)
	fmt.Println("ServerCfg cfgFile:", ServerCfg.CfgFile)
	fmt.Println("ServerCfg host:", ServerCfg.Host)

	if ServerCfg.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ServerCfg.CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("demo")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
