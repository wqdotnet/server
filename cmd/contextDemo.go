package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// contextDemoCmd represents the contextDemo command
var contextDemoCmd = &cobra.Command{
	Use:   "contextDemo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		demo()
	},
}

func init() {
	rootCmd.AddCommand(contextDemoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contextDemoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contextDemoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var c = 1

func doSome(i int) error {
	c++
	fmt.Println(c)
	if c > 3 {
		return errors.New("err occur")
	}
	return nil
}

func speakMemo(ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Done")
			return
		default:
			fmt.Println("exec default func")
			err := doSome(3)
			if err != nil {
				fmt.Printf("cancelFunc()")
				cancelFunc()
			}
		}
	}
}

func demo() {
	rootContext := context.Background()
	ctx, cancelFunc := context.WithCancel(rootContext)
	go speakMemo(ctx, cancelFunc)
	time.Sleep(time.Second * 5)
}
