package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// ServerConfig  server cfg
type ServerConfig struct {
	OpenHTTP string
	HTTPPort int

	NetWork string
	Port    int

	//proto_path=%s  --go_out

	ProtoPath string
	GoOut     string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{

	OpenHTTP: "localhost",
	HTTPPort: 8080,

	// #network : tcp/udp
	NetWork: "tcp",
	Port:    3344,

	// #protobuf path
	ProtoPath: "E:/gopath/src/server/proto",
	GoOut:     "E:/gopath/src/server/proto",
}

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

	rootCmd.PersistentFlags().StringVar(&cfgfile, "config", "", "config file (default is $HOME/demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgfile != "" {
		fmt.Println("initconfig ", cfgfile)
		// Use config file from the flag.
		viper.SetConfigFile(cfgfile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("initconfig ", home)

		// Search config in home directory with name ".demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("cfg")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
