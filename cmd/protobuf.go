package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// protobufCmd build protobuf
var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "Short",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {

		//protoc --proto_path=d:/proto  --go_out=d:/proto  msg.pro

		out, errout, err := Shellout(args...)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if out != "" {
			fmt.Println("--- stdout ---")
			fmt.Println(out)
		}
		if errout != "" {
			fmt.Println("--- stderr ---")
			fmt.Println(errout)
		}
	},
}

func init() {
	rootCmd.AddCommand(protobufCmd)

}
